package domain

import (
	"encoding/json"
	"fmt"
)

// BocErrorCode ...
var BocErrorCode map[string]int

type (

	// BocCacheType ...
	BocCacheType struct {
		ValueEnumType interface{}
	}

	// BocCacheTypePinned ...
	BocCacheTypePinned struct {
		Pin string `json:"pin"`
	}

	// BocCacheTypeUnpinned ...
	BocCacheTypeUnpinned struct{}

	// ParamsOfParse ...
	ParamsOfParse struct {
		Boc string `json:"boc"`
	}

	// ResultOfParse ...
	ResultOfParse struct {
		Parsed json.RawMessage `json:"parsed"`
	}

	// ParamsOfParseShardstate ...
	ParamsOfParseShardstate struct {
		Boc         string `json:"boc"`
		ID          string `json:"id"`
		WorkchainID int    `json:"workchain_id"`
	}

	// ParamsOfGetBlockchainConfig ...
	ParamsOfGetBlockchainConfig struct {
		BlockBoc string `json:"block_boc"`
	}

	// ResultOfGetBlockchainConfig ...
	ResultOfGetBlockchainConfig struct {
		ConfigBoc string `json:"config_boc"`
	}

	// ParamsOfGetBocHash ...
	ParamsOfGetBocHash struct {
		Boc string `json:"boc"`
	}

	// ResultOfGetBocHash ...
	ResultOfGetBocHash struct {
		Hash string `json:"hash"`
	}

	//ParamsOfGetCodeFromTvc ...
	ParamsOfGetCodeFromTvc struct {
		Tvc string `json:"tvc"`
	}

	//ResultOfGetCodeFromTvc ...
	ResultOfGetCodeFromTvc struct {
		Code string `json:"code"`
	}

	// ParamsOfBocCacheGet ...
	ParamsOfBocCacheGet struct {
		BocRef string `json:"boc_ref"`
	}

	// ResultOfBocCacheGet ...
	ResultOfBocCacheGet struct {
		Boc string `json:"boc,omitempty"`
	}

	// ParamsOfBocCacheSet ...
	ParamsOfBocCacheSet struct {
		Boc       string        `json:"boc"`
		CacheType *BocCacheType `json:"cache_type"`
	}

	// ResultOfBocCacheSet ...
	ResultOfBocCacheSet struct {
		BocRef string `json:"boc_ref"`
	}

	// ParamsOfBocCacheUnpin ...
	ParamsOfBocCacheUnpin struct {
		Pin    string `json:"pin"`
		BocRef string `json:"boc_ref,omitempty"`
	}

	// BuilderOp - Cell builder operation.
	BuilderOp struct {
		ValueEnumType interface{}
	}

	// BuilderOpInteger ...
	BuilderOpInteger struct {
		Size  int         `json:"size"`
		Value interface{} `json:"value"`
	}

	// BuilderOpBitString ...
	BuilderOpBitString struct {
		Value string `json:"value"`
	}

	// BuilderOpCell ...
	BuilderOpCell struct {
		Builder []*BuilderOp `json:"builder"`
	}

	// BuilderOpCellBoc ...
	BuilderOpCellBoc struct {
		Boc string `json:"boc"`
	}

	// ParamsOfEncodeBoc ...
	ParamsOfEncodeBoc struct {
		Builder  []*BuilderOp  `json:"builder"`
		BocCache *BocCacheType `json:"boc_cache,omitempty"`
	}
	// ResultOfEncodeBoc ...
	ResultOfEncodeBoc struct {
		Boc string
	}

	//BocUseCase ...
	BocUseCase interface {
		ParseMessage(*ParamsOfParse) (*ResultOfParse, error)
		ParseTransaction(*ParamsOfParse) (*ResultOfParse, error)
		ParseAccount(*ParamsOfParse) (*ResultOfParse, error)
		ParseBlock(*ParamsOfParse) (*ResultOfParse, error)
		ParseShardstate(*ParamsOfParseShardstate) (*ResultOfParse, error)
		GetBlockhainConfig(*ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error)
		GetBocHash(*ParamsOfGetBocHash) (*ResultOfGetBocHash, error)
		GetCodeFromTvc(*ParamsOfGetCodeFromTvc) (*ResultOfGetCodeFromTvc, error)
		CacheGet(*ParamsOfBocCacheGet) (*ResultOfBocCacheGet, error)
		CacheSet(*ParamsOfBocCacheSet) (*ResultOfBocCacheSet, error)
		CacheUnpin(*ParamsOfBocCacheUnpin) error
		EncodeBoc(*ParamsOfEncodeBoc) (*ResultOfEncodeBoc, error)
	}
)

func init() {
	BocErrorCode = map[string]int{
		"InvalidBoc":            201,
		"SerializationError":    202,
		"InappropriateBlock":    203,
		"MissingSourceBoc":      204,
		"InsufficientCacheSize": 205,
		"BocRefNotFound":        206,
		"InvalidBocRef":         207,
	}
}

func (bCT *BocCacheType) MarshalJSON() ([]byte, error) {
	switch value := (bCT.ValueEnumType).(type) {
	case BocCacheTypePinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			BocCacheTypePinned
		}{"Pinned", value})
	case BocCacheTypeUnpinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			BocCacheTypeUnpinned
		}{"Unpinned", value})
	default:
		return nil, fmt.Errorf("unsupported type for BocCacheType %v", bCT.ValueEnumType)
	}
}

func (bCT *BocCacheType) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Pinned":
		var valueEnum BocCacheTypePinned
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bCT.ValueEnumType = valueEnum
	case "Unpinned":
		var valueEnum BocCacheTypeUnpinned
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bCT.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for BocCacheType %v", typeD.Type)
	}
	return nil
}

// NewBocCacheType ...
func NewBocCacheType(value interface{}) *BocCacheType {
	return &BocCacheType{ValueEnumType: value}
}

func (bO *BuilderOp) MarshalJSON() ([]byte, error) {
	switch value := (bO.ValueEnumType).(type) {
	case BuilderOpInteger:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpInteger
		}{"Integer", value})
	case BuilderOpBitString:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpBitString
		}{"BitString", value})
	case BuilderOpCell:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpCell
		}{"Cell", value})
	case BuilderOpCellBoc:
		return json.Marshal(struct {
			Type string `json:"type"`
			BuilderOpCellBoc
		}{"CellBoc", value})
	default:
		return nil, fmt.Errorf("unsupported type for BuilderOp %v", bO.ValueEnumType)
	}
}

func (bO *BuilderOp) UnmarshalJSON(b []byte) error {
	var typeD EnumType
	if err := json.Unmarshal(b, &typeD); err != nil {
		return err
	}

	switch typeD.Type {
	case "Integer":
		var valueEnum BuilderOpInteger
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "BitString":
		var valueEnum BuilderOpBitString
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "Cell":
		var valueEnum BuilderOpCell
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	case "CellBoc":
		var valueEnum BuilderOpCellBoc
		if err := json.Unmarshal(b, &valueEnum); err != nil {
			return err
		}
		bO.ValueEnumType = valueEnum
	default:
		return fmt.Errorf("unsupported type for BuilderOp %v", typeD.Type)
	}
	return nil
}

// NewBuilderOp ...
func NewBuilderOp(value interface{}) *BuilderOp {
	return &BuilderOp{ValueEnumType: value}
}
