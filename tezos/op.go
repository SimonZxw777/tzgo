// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tezos

import (
	"fmt"
)

type OpStatus byte

const (
	OpStatusInvalid OpStatus = iota // 0
	OpStatusApplied                 // 1 (success)
	OpStatusFailed
	OpStatusSkipped
	OpStatusBacktracked
)

func (t OpStatus) IsValid() bool {
	return t != OpStatusInvalid
}

func (t OpStatus) IsSuccess() bool {
	return t == OpStatusApplied
}

func (t *OpStatus) UnmarshalText(data []byte) error {
	v := ParseOpStatus(string(data))
	if !v.IsValid() {
		return fmt.Errorf("tezos: invalid operation status '%s'", string(data))
	}
	*t = v
	return nil
}

func (t OpStatus) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func ParseOpStatus(s string) OpStatus {
	switch s {
	case "applied":
		return OpStatusApplied
	case "failed":
		return OpStatusFailed
	case "skipped":
		return OpStatusSkipped
	case "backtracked":
		return OpStatusBacktracked
	default:
		return OpStatusInvalid
	}
}

func (t OpStatus) String() string {
	switch t {
	case OpStatusApplied:
		return "applied"
	case OpStatusFailed:
		return "failed"
	case OpStatusSkipped:
		return "skipped"
	case OpStatusBacktracked:
		return "backtracked"
	default:
		return ""
	}
}

type OpType byte

// enums are allocated in chronological order
const (
	OpTypeInvalid                      OpType = iota
	OpTypeActivateAccount                     // 1
	OpTypeDoubleBakingEvidence                // 2
	OpTypeDoubleEndorsementEvidence           // 3
	OpTypeSeedNonceRevelation                 // 4
	OpTypeTransaction                         // 5
	OpTypeOrigination                         // 6
	OpTypeDelegation                          // 7
	OpTypeReveal                              // 8
	OpTypeEndorsement                         // 9
	OpTypeProposals                           // 10
	OpTypeBallot                              // 11
	OpTypeFailingNoop                         // 12 v009
	OpTypeEndorsementWithSlot                 // 13 v009
	OpTypeRegisterConstant                    // 14 v011
	OpTypePreendorsement                      // 15 v012
	OpTypeDoublePreendorsementEvidence        // 16 v012
	OpTypeSetDepositsLimit                    // 17 v012
	OpTypeToruOrigination                     // 18 v013
	OpTypeToruSubmitBatch                     // 19 v013
	OpTypeToruCommit                          // 20 v013
	OpTypeToruReturnBond                      // 21 v013
	OpTypeToruFinalizeCommitment              // 22 v013
	OpTypeToruRemoveCommitment                // 23 v013
	OpTypeToruRejection                       // 24 v013
	OpTypeToruDispatchTickets                 // 25 v013
	OpTypeTransferTicket                      // 26 v013
	OpTypeScRollupOriginate                   // 27 v013
	OpTypeScRollupAddMessages                 // 28 v013
	OpTypeScRollupCement                      // 29 v013
	OpTypeScRollupPublish                     // 30 v013
	OpTypeVdfRevelation                       // 31 v014
	OpTypeIncreasePaidStorage                 // 32 v014
	OpTypeScRollupRefute                      // 33 v014
	OpTypeScRollupTimeout                     // 34 v014
	OpTypeScRollupExecuteOutboxMessage        // 35 v014
	OpTypeScRollupRecoverBond                 // 36 v014
	OpTypeScRollupDalSlotSubscribe            // 37 v014
	OpTypeDalSlotAvailability                 // 38 v014
	OpTypeDalPublishSlotHeader                // 39 v014
	OpTypeEvent                               // 40 v014 (only in internal_operation_results)
	OpTypeDrainDelegate                       // 41 v015
	OpTypeUpdateConsensusKey                  // 42 v015
)

var (
	opTypeStrings = map[OpType]string{
		OpTypeInvalid:                      "",
		OpTypeActivateAccount:              "activate_account",
		OpTypeDoubleBakingEvidence:         "double_baking_evidence",
		OpTypeDoubleEndorsementEvidence:    "double_endorsement_evidence",
		OpTypeSeedNonceRevelation:          "seed_nonce_revelation",
		OpTypeTransaction:                  "transaction",
		OpTypeOrigination:                  "origination",
		OpTypeDelegation:                   "delegation",
		OpTypeReveal:                       "reveal",
		OpTypeEndorsement:                  "endorsement",
		OpTypeProposals:                    "proposals",
		OpTypeBallot:                       "ballot",
		OpTypeFailingNoop:                  "failing_noop",
		OpTypeEndorsementWithSlot:          "endorsement_with_slot",
		OpTypeRegisterConstant:             "register_global_constant",
		OpTypePreendorsement:               "preendorsement",
		OpTypeDoublePreendorsementEvidence: "double_preendorsement_evidence",
		OpTypeSetDepositsLimit:             "set_deposits_limit",
		OpTypeToruOrigination:              "tx_rollup_origination",
		OpTypeToruSubmitBatch:              "tx_rollup_submit_batch",
		OpTypeToruCommit:                   "tx_rollup_commit",
		OpTypeToruReturnBond:               "tx_rollup_return_bond",
		OpTypeToruFinalizeCommitment:       "tx_rollup_finalize_commitment",
		OpTypeToruRemoveCommitment:         "tx_rollup_remove_commitment",
		OpTypeToruRejection:                "tx_rollup_rejection",
		OpTypeToruDispatchTickets:          "tx_rollup_dispatch_tickets",
		OpTypeTransferTicket:               "transfer_ticket",
		OpTypeScRollupOriginate:            "sc_rollup_originate",
		OpTypeScRollupAddMessages:          "sc_rollup_add_messages",
		OpTypeScRollupCement:               "sc_rollup_cement",
		OpTypeScRollupPublish:              "sc_rollup_publish",
		OpTypeVdfRevelation:                "vdf_revelation",
		OpTypeIncreasePaidStorage:          "increase_paid_storage",
		OpTypeScRollupRefute:               "sc_rollup_refute",
		OpTypeScRollupTimeout:              "sc_rollup_timeout",
		OpTypeScRollupExecuteOutboxMessage: "sc_rollup_execute_outbox_message",
		OpTypeScRollupRecoverBond:          "sc_rollup_recover_bond",
		OpTypeScRollupDalSlotSubscribe:     "sc_rollup_dal_slot_subscribe",
		OpTypeDalSlotAvailability:          "dal_slot_availability",
		OpTypeDalPublishSlotHeader:         "dal_publish_slot_header",
		OpTypeEvent:                        "event",
		OpTypeDrainDelegate:                "drain_delegate",
		OpTypeUpdateConsensusKey:           "update_consensus_key",
	}
	opTypeReverseStrings = make(map[string]OpType)
)

func init() {
	for n, v := range opTypeStrings {
		opTypeReverseStrings[v] = n
	}
}

func (t OpType) IsValid() bool {
	return t != OpTypeInvalid
}

func (t *OpType) UnmarshalText(data []byte) error {
	v := ParseOpType(string(data))
	if !v.IsValid() {
		return fmt.Errorf("tezos: invalid operation type '%s'", string(data))
	}
	*t = v
	return nil
}

func (t OpType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func ParseOpType(s string) OpType {
	t, ok := opTypeReverseStrings[s]
	if !ok {
		t = OpTypeInvalid
	}
	return t
}

func (t OpType) String() string {
	return opTypeStrings[t]
}

var (
	// before Babylon v005
	opTagV0 = map[OpType]byte{
		OpTypeEndorsement:               0,
		OpTypeSeedNonceRevelation:       1,
		OpTypeDoubleEndorsementEvidence: 2,
		OpTypeDoubleBakingEvidence:      3,
		OpTypeActivateAccount:           4,
		OpTypeProposals:                 5,
		OpTypeBallot:                    6,
		OpTypeReveal:                    7,
		OpTypeTransaction:               8,
		OpTypeOrigination:               9,
		OpTypeDelegation:                10,
	}
	// Babylon v005 until Hangzhou v011
	opTagV1 = map[OpType]byte{
		OpTypeEndorsement:               0,
		OpTypeSeedNonceRevelation:       1,
		OpTypeDoubleEndorsementEvidence: 2,
		OpTypeDoubleBakingEvidence:      3,
		OpTypeActivateAccount:           4,
		OpTypeProposals:                 5,
		OpTypeBallot:                    6,
		OpTypeReveal:                    107, // v005
		OpTypeTransaction:               108, // v005
		OpTypeOrigination:               109, // v005
		OpTypeDelegation:                110, // v005
		OpTypeEndorsementWithSlot:       10,  // v009
		OpTypeFailingNoop:               17,  // v009
		OpTypeRegisterConstant:          111, // v011
	}
	// Ithaca v012 and up
	opTagV2 = map[OpType]byte{
		OpTypeSeedNonceRevelation:          1,
		OpTypeDoubleEndorsementEvidence:    2,
		OpTypeDoubleBakingEvidence:         3,
		OpTypeActivateAccount:              4,
		OpTypeProposals:                    5,
		OpTypeBallot:                       6,
		OpTypeReveal:                       107, // v005
		OpTypeTransaction:                  108, // v005
		OpTypeOrigination:                  109, // v005
		OpTypeDelegation:                   110, // v005
		OpTypeFailingNoop:                  17,  // v009
		OpTypeRegisterConstant:             111, // v011
		OpTypePreendorsement:               20,  // v012
		OpTypeEndorsement:                  21,  // v012
		OpTypeDoublePreendorsementEvidence: 7,   // v012
		OpTypeSetDepositsLimit:             112, // v012
		OpTypeToruOrigination:              150, // v013
		OpTypeToruSubmitBatch:              151, // v013
		OpTypeToruCommit:                   152, // v013
		OpTypeToruReturnBond:               153, // v013
		OpTypeToruFinalizeCommitment:       154, // v013
		OpTypeToruRemoveCommitment:         155, // v013
		OpTypeToruRejection:                156, // v013
		OpTypeToruDispatchTickets:          157, // v013
		OpTypeTransferTicket:               158, // v013
		OpTypeScRollupOriginate:            200, // v013
		OpTypeScRollupAddMessages:          201, // v013
		OpTypeScRollupCement:               202, // v013
		OpTypeScRollupPublish:              203, // v013
		OpTypeVdfRevelation:                8,   // v014
		OpTypeDalSlotAvailability:          22,  // v014
		OpTypeScRollupRefute:               204, // v014
		OpTypeScRollupTimeout:              205, // v014
		OpTypeScRollupExecuteOutboxMessage: 206, // v014
		OpTypeScRollupRecoverBond:          207, // v014
		OpTypeScRollupDalSlotSubscribe:     208, // v014
		OpTypeDalPublishSlotHeader:         230, // v014
		OpTypeDrainDelegate:                9,   // v015
		OpTypeUpdateConsensusKey:           114, // v015
	}
)

func (t OpType) TagVersion(ver int) byte {
	var (
		tag byte
		ok  bool
	)
	switch ver {
	case 0:
		tag, ok = opTagV0[t]
	case 1:
		tag, ok = opTagV1[t]
	default:
		tag, ok = opTagV2[t]
	}
	if !ok {
		return 255
	}
	return tag
}

func (t OpType) Tag() byte {
	tag, ok := opTagV2[t]
	if !ok {
		tag = 255
	}
	return tag
}

var (
	// before Babylon v005
	opMinSizeV0 = map[byte]int{
		0:  5,         // OpTypeEndorsement
		1:  37,        // OpTypeSeedNonceRevelation
		2:  9 + 2*101, // OpTypeDoubleEndorsementEvidence
		3:  9 + 2*189, // OpTypeDoubleBakingEvidence (w/o seed_nonce_hash)
		4:  41,        // OpTypeActivateAccount
		5:  30,        // OpTypeProposals
		6:  59,        // OpTypeBallot
		7:  26 + 32,   // OpTypeReveal (assuming shortest pk)
		8:  49,        // OpTypeTransaction
		9:  53,        // OpTypeOrigination
		10: 26,        // OpTypeDelegation
	}
	// Babylon v005 until Hangzhou v011
	opMinSizeV1 = map[byte]int{
		0:   5,          // OpTypeEndorsement // <v009
		1:   37,         // OpTypeSeedNonceRevelation
		2:   11 + 2*101, // OpTypeDoubleEndorsementEvidence
		3:   9 + 2*189,  // OpTypeDoubleBakingEvidence (w/o seed_nonce_hash, lb_escape_vote)
		4:   41,         // OpTypeActivateAccount
		5:   30,         // OpTypeProposals
		6:   59,         // OpTypeBallot
		107: 26 + 32,    // OpTypeReveal // v005 (assuming shortest pk)
		108: 50,         // OpTypeTransaction // v005
		109: 28,         // OpTypeOrigination // v005
		110: 27,         // OpTypeDelegation // v005
		10:  108,        // OpTypeEndorsementWithSlot // v009
		17:  5,          // OpTypeFailingNoop  // v009
		111: 30,         // OpTypeRegisterConstant // v011
	}
	// Ithaca v012 and up
	opMinSizeV2 = map[byte]int{
		1:   37,               // OpTypeSeedNonceRevelation
		2:   9 + 2*(32+43+64), // OpTypeDoubleEndorsementEvidence
		3:   9 + 2*237,        // OpTypeDoubleBakingEvidence (w/o seed_nonce_hash, min fitness size)
		4:   41,               // OpTypeActivateAccount
		5:   30,               // OpTypeProposals
		6:   59,               // OpTypeBallot
		107: 26 + 32,          // OpTypeReveal // v005 (assuming shortest pk)
		108: 50,               // OpTypeTransaction // v005
		109: 28,               // OpTypeOrigination // v005
		110: 27,               // OpTypeDelegation // v005
		17:  5,                // OpTypeFailingNoop  // v009
		111: 30,               // OpTypeRegisterConstant // v011
		7:   9 + 2*(32+43+64), // OpTypeDoublePreendorsementEvidence // v012
		20:  43,               // OpTypePreendorsement // v012
		21:  43,               // OpTypeEndorsement // v012
		112: 27,               // OpTypeSetDepositsLimit // v012
		9:   1 + 3*21,         // OpTypeDrainDelegate // v015
		114: 26 + 32,          // OpTypeUpdateConsensusKey // v015
	}
)

func (t OpType) MinSizeVersion(ver int) int {
	switch ver {
	case 0:
		return opMinSizeV0[t.TagVersion(ver)]
	case 1:
		return opMinSizeV1[t.TagVersion(ver)]
	default:
		return opMinSizeV2[t.TagVersion(ver)]
	}
}

func (t OpType) MinSize() int {
	return opMinSizeV2[t.Tag()]
}

func (t OpType) ListId() int {
	switch t {
	case OpTypeEndorsement, OpTypeEndorsementWithSlot, OpTypePreendorsement:
		return 0
	case OpTypeProposals, OpTypeBallot:
		return 1
	case OpTypeActivateAccount,
		OpTypeDoubleBakingEvidence,
		OpTypeDoubleEndorsementEvidence,
		OpTypeSeedNonceRevelation,
		OpTypeDoublePreendorsementEvidence,
		OpTypeVdfRevelation,
		OpTypeDrainDelegate:
		return 2
	case OpTypeTransaction, // generic user operations
		OpTypeOrigination,
		OpTypeDelegation,
		OpTypeReveal,
		OpTypeRegisterConstant,
		OpTypeSetDepositsLimit,
		OpTypeToruOrigination,
		OpTypeToruSubmitBatch,
		OpTypeToruCommit,
		OpTypeToruReturnBond,
		OpTypeToruFinalizeCommitment,
		OpTypeToruRemoveCommitment,
		OpTypeToruRejection,
		OpTypeToruDispatchTickets,
		OpTypeTransferTicket,
		OpTypeScRollupOriginate,
		OpTypeScRollupAddMessages,
		OpTypeScRollupCement,
		OpTypeScRollupPublish,
		OpTypeDalSlotAvailability,
		OpTypeScRollupRefute,
		OpTypeScRollupTimeout,
		OpTypeScRollupExecuteOutboxMessage,
		OpTypeScRollupRecoverBond,
		OpTypeScRollupDalSlotSubscribe,
		OpTypeDalPublishSlotHeader,
		OpTypeUpdateConsensusKey:
		return 3
	default:
		return -1 // invalid
	}
}

func ParseOpTag(t byte) OpType {
	switch t {
	case 0:
		return OpTypeEndorsement
	case 10:
		return OpTypeEndorsementWithSlot
	case 1:
		return OpTypeSeedNonceRevelation
	case 2:
		return OpTypeDoubleEndorsementEvidence
	case 3:
		return OpTypeDoubleBakingEvidence
	case 4:
		return OpTypeActivateAccount
	case 5:
		return OpTypeProposals
	case 6:
		return OpTypeBallot
	case 107:
		return OpTypeReveal
	case 108:
		return OpTypeTransaction
	case 109:
		return OpTypeOrigination
	case 110:
		return OpTypeDelegation
	case 17:
		return OpTypeFailingNoop
	case 111:
		return OpTypeRegisterConstant
	case 20:
		return OpTypePreendorsement
	case 21:
		return OpTypeEndorsement
	case 7:
		return OpTypeDoublePreendorsementEvidence
	case 112:
		return OpTypeSetDepositsLimit
	case 150:
		return OpTypeToruOrigination
	case 151:
		return OpTypeToruSubmitBatch
	case 152:
		return OpTypeToruCommit
	case 153:
		return OpTypeToruReturnBond
	case 154:
		return OpTypeToruFinalizeCommitment
	case 155:
		return OpTypeToruRemoveCommitment
	case 156:
		return OpTypeToruRejection
	case 157:
		return OpTypeToruDispatchTickets
	case 158:
		return OpTypeTransferTicket
	case 200:
		return OpTypeScRollupOriginate
	case 201:
		return OpTypeScRollupAddMessages
	case 202:
		return OpTypeScRollupCement
	case 203:
		return OpTypeScRollupPublish
	case 8:
		return OpTypeVdfRevelation
	case 22:
		return OpTypeDalSlotAvailability
	case 204:
		return OpTypeScRollupRefute
	case 205:
		return OpTypeScRollupTimeout
	case 206:
		return OpTypeScRollupExecuteOutboxMessage
	case 207:
		return OpTypeScRollupRecoverBond
	case 208:
		return OpTypeScRollupDalSlotSubscribe
	case 230:
		return OpTypeDalPublishSlotHeader
	case 9:
		return OpTypeDrainDelegate
	case 114:
		return OpTypeUpdateConsensusKey
	default:
		return OpTypeInvalid
	}
}

func ParseOpTagVersion(t byte, ver int) OpType {
	tags := opTagV0
	switch ver {
	case 1:
		tags = opTagV1
	case 2:
		tags = opTagV2
	}
	for typ, tag := range tags {
		if tag == t {
			return typ
		}
	}
	return OpTypeInvalid
}
