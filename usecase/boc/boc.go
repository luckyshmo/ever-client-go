package boc

import "github.com/move-ton/ever-client-go/domain"

type boc struct {
	config domain.Config
	client domain.ClientGateway
}

// NewBoc ...
func NewBoc(
	config domain.Config,
	client domain.ClientGateway,
) domain.BocUseCase {
	return &boc{
		config: config,
		client: client,
	}
}

// ParseMessage - Parses message boc into a JSON.
func (b *boc) ParseMessage(pOP *domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_message", pOP, result)
	return result, err
}

// ParseTransaction - Parses transaction boc into a JSON.
func (b *boc) ParseTransaction(pOP *domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_transaction", pOP, result)
	return result, err
}

// ParseAccount - Parses account boc into a JSON.
func (b *boc) ParseAccount(pOP *domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_account", pOP, result)
	return result, err
}

// ParseBlock - Parses block boc into a JSON.
func (b *boc) ParseBlock(pOP *domain.ParamsOfParse) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_block", pOP, result)
	return result, err
}

// ParseShardstate - Parses shardstate boc into a JSON.
func (b *boc) ParseShardstate(pOPS *domain.ParamsOfParseShardstate) (*domain.ResultOfParse, error) {
	result := new(domain.ResultOfParse)
	err := b.client.GetResult("boc.parse_shardstate", pOPS, result)
	return result, err
}

// GetBlockhainConfig - Extract blockchain configuration from key block and also from zerostate.
func (b *boc) GetBlockhainConfig(pOGBC *domain.ParamsOfGetBlockchainConfig) (*domain.ResultOfGetBlockchainConfig, error) {
	result := new(domain.ResultOfGetBlockchainConfig)
	err := b.client.GetResult("boc.get_blockchain_config", pOGBC, result)
	return result, err
}

// GetBocHash - Calculates BOC root hash.
func (b *boc) GetBocHash(pOGBH *domain.ParamsOfGetBocHash) (*domain.ResultOfGetBocHash, error) {
	result := new(domain.ResultOfGetBocHash)
	err := b.client.GetResult("boc.get_boc_hash", pOGBH, result)
	return result, err
}

// GetBocDepth - Calculates BOC depth.
func (b *boc) GetBocDepth(pOGBD *domain.ParamsOfGetBocDepth) (*domain.ResultOfGetBocDepth, error) {
	result := new(domain.ResultOfGetBocDepth)
	err := b.client.GetResult("boc.get_boc_depth", pOGBD, result)
	return result, err
}

// GetCodeFromTvc - Extracts code from TVC contract image.
func (b *boc) GetCodeFromTvc(pOGCFT *domain.ParamsOfGetCodeFromTvc) (*domain.ResultOfGetCodeFromTvc, error) {
	result := new(domain.ResultOfGetCodeFromTvc)
	err := b.client.GetResult("boc.get_code_from_tvc", pOGCFT, result)
	return result, err
}

// CacheGet - Get BOC from cache.
func (b *boc) CacheGet(pOBCG *domain.ParamsOfBocCacheGet) (*domain.ResultOfBocCacheGet, error) {
	result := new(domain.ResultOfBocCacheGet)
	err := b.client.GetResult("boc.cache_get", pOBCG, result)
	return result, err
}

// CacheSet - Save BOC into cache.
func (b *boc) CacheSet(pOBCS *domain.ParamsOfBocCacheSet) (*domain.ResultOfBocCacheSet, error) {
	result := new(domain.ResultOfBocCacheSet)
	err := b.client.GetResult("boc.cache_set", pOBCS, result)
	return result, err
}

// CacheUnpin - Unpin BOCs with specified pin.
// BOCs which don't have another pins will be removed from cache.
func (b *boc) CacheUnpin(pOBCU *domain.ParamsOfBocCacheUnpin) error {
	_, err := b.client.GetResponse("boc.cache_unpin", pOBCU)
	return err
}

// EncodeBoc - Encodes BOC from builder operations.
func (b *boc) EncodeBoc(pOEB *domain.ParamsOfEncodeBoc) (*domain.ResultOfEncodeBoc, error) {
	result := new(domain.ResultOfEncodeBoc)
	err := b.client.GetResult("boc.encode_boc", pOEB, result)
	return result, err
}

// GetCodeSalt - Returns the contract code's salt if it is present.
func (b *boc) GetCodeSalt(pOGCS *domain.ParamsOfGetCodeSalt) (*domain.ResultOfGetCodeSalt, error) {
	result := new(domain.ResultOfGetCodeSalt)
	err := b.client.GetResult("boc.get_code_salt", pOGCS, result)
	return result, err
}

// SetCodeSalt - Sets new salt to contract code.
// Returns the new contract code with salt.
func (b *boc) SetCodeSalt(pOSCS *domain.ParamsOfSetCodeSalt) (*domain.ResultOfSetCodeSalt, error) {
	result := new(domain.ResultOfSetCodeSalt)
	err := b.client.GetResult("boc.set_code_salt", pOSCS, result)
	return result, err
}

// DecodeTvc - Decodes tvc into code, data, libraries and special options.
func (b *boc) DecodeTvc(pODT *domain.ParamsOfDecodeTvc) (*domain.ResultOfDecodeTvc, error) {
	result := new(domain.ResultOfDecodeTvc)
	err := b.client.GetResult("boc.decode_tvc", pODT, result)
	return result, err
}

// EncodeTvc - Encodes tvc from code, data, libraries ans special options (see input params).
func (b *boc) EncodeTvc(pOET *domain.ParamsOfEncodeTvc) (*domain.ResultOfEncodeTvc, error) {
	result := new(domain.ResultOfEncodeTvc)
	err := b.client.GetResult("boc.encode_tvc", pOET, result)
	return result, err
}

// GetCompilerVersion - Returns the compiler version used to compile the code.
func (b *boc) GetCompilerVersion(pOGCV *domain.ParamsOfGetCompilerVersion) (*domain.ResultOfGetCompilerVersion, error) {
	result := new(domain.ResultOfGetCompilerVersion)
	err := b.client.GetResult("boc.get_compiler_version", pOGCV, result)
	return result, err
}
