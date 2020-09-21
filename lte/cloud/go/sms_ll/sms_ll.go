package sms_ll

import (
	"github.com/warthog618/sms"
	"github.com/warthog618/sms/encoding/tpdu"
	"time"
)

// Generate fully encoded SMS PDUs for delivery to a UE (MS). Will handle
// encoding and chunking of messages as appropriate. We first generate TPDUs,
// then RP-DATA headers, and finally CP-DATA headers, resulting in a set of
// byte arrays that can be directly delivered to a UE (MS).
// Inputs:
// 	message: A UTF-8 string representing the SMS.
//	from_num: A E.164 encoded source number.
//	timestamp: The sender timestamp for the SMS (generally, use current server time)
//	ref_base: The starting reference number for this set of SMS messages. Should be a counter per IMSI.
// Outputs:
//	- Array of byte array representing the set of fully-encoded CP-DATA(RP-DATA(TPDU)) messages generated
//	- Error	(if any)
func GenerateSMSDelivers(message string, from_num string, timestamp time.Time, ref_base uint8) ([][]byte, error) {
	tpdus := createTPDUs(message, from_num, timestamp)

	output := make([][]byte, 0)

	for i := range tpdus {
		tp, err := tpdus[i].MarshalBinary()
		if err != nil {
			return nil, err
		}
		reference := byte(int(ref_base) + i)
		rpm, err := createRPDataMessage(RP_MTI_MT_DATA, reference, tp)
		if err != nil {
			return nil, err
		}

		marshaledRPM, err := rpm.MarshalBinary()
		if err != nil {
			return nil, err
		}

		cpm, err := createCPDataMessage(marshaledRPM, reference)
		if err != nil {
			return nil, err
		}

		marshaledCPM, err := cpm.MarshalBinary()
		if err != nil {
			return nil, err
		}

		output = append(output, marshaledCPM)
	}

	return output, nil
}

func createTPDUs(message string, from_num string, timestamp time.Time) []tpdu.TPDU {
	tpdus, _ := sms.Encode([]byte(message), sms.AsDeliver, sms.From(from_num))
	for i := range tpdus {
		tpdus[i].FirstOctet |= tpdu.FoMMS // Android won't accept if this bit isn't set.
		tpdus[i].FirstOctet |= tpdu.FoSRI // Request a delivery report.
		tpdus[i].SCTS = tpdu.Timestamp{Time: timestamp}
	}

	return tpdus
}

// Helper for creating a RP-DATA message. If is_mt is true, make this mobile terminated.
func createRPDataMessage(mti byte, reference byte, data []byte) (*RPMessage, error) {
	rpm := new(RPMessage)
	rpm.MTI = mti
	rpm.Reference = reference

	msg_type, err := rpm.MsgType()
	if msg_type != RP_DATA {
		return nil, SMSRPError("MTI isn't RP-DATA")
	} else if err != nil {
		return nil, err
	}

	if rpm.Direction() == RP_MT { // MT-SMS should have empty dest address
		rpm.OriginatorAddress, rpm.DestinationAddress = NewFakeRPAddressElement(), RPAddressElement{Length: 0}
	} else { // MO-SMS should have empty orig address
		rpm.OriginatorAddress, rpm.DestinationAddress = RPAddressElement{Length: 0}, NewFakeRPAddressElement()
	}

	rpm.UserData, err = createRPUserElement(data)
	if err != nil {
		return nil, err
	}
	return rpm, nil
}

func createCPDataMessage(rpdu []byte, txID byte) (CPMessage, error) {
	cpm, err := CreateCPMessage(txID, CP_DATA, rpdu)
	if err != nil {
		return cpm, err
	}
	return cpm, nil
}
