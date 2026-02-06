package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Salvionied/apollo"
	"github.com/Salvionied/apollo/constants"
	"github.com/Salvionied/apollo/txBuilding/Backend/BlockFrostChainContext"
	"github.com/fxamacker/cbor/v2"
)

func main() {
    bfc, err := BlockFrostChainContext.NewBlockfrostChainContext(constants.BLOCKFROST_BASE_URL_PREPROD, int(constants.PREPROD), "preprod01ZVfnDyuVMUaTsRqhoT0njyYD1fZ12L")
    if err != nil {
        panic(err)
    }

    cc := apollo.NewEmptyBackend()
    SEED := "silent scheme maple double kiwi sleep jump embody obscure omit brand flee immune toe minor increase code network street grid allow protect gift spray"
    apollob := apollo.New(&cc)
    apollob, err = apollob.SetWalletFromMnemonic(SEED, constants.PREPROD)
    if err != nil {
        panic(err)
    }
    apollob, err = apollob.SetWalletAsChangeAddress()
    if err != nil {
        panic(err)
    }
	fmt.Println(apollob.GetWallet().GetAddress())
    utxos, err := bfc.Utxos(*apollob.GetWallet().GetAddress())
    if err != nil {
        panic(err)
    }
    apollob, err = apollob.AddLoadedUTxOs(utxos...).PayToAddressBech32("addr_test1qza22lw5zl782d7pn3w3c7m2342kwdmlamjcfw9znmh5v5prts2g3cjle2luxam9rayfxq5s962ze8r8mjauxahxawuste3dks", 1_000_000).
        Complete()
    if err != nil {
        panic(err)
    }
    apollob = apollob.Sign()
    tx := apollob.GetTx()
    cborred, err := cbor.Marshal(tx)
    if err != nil {
        panic(err)
    }
    fmt.Println(hex.EncodeToString(cborred))
    tx_id, err := bfc.SubmitTx(*tx)
    if err != nil {
		panic(err)
	}
    fmt.Println(hex.EncodeToString(tx_id.Payload))

}
