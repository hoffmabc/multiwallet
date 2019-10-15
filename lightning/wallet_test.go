package lightning

import (
	"crypto/rand"
	"github.com/OpenBazaar/multiwallet/datastore"
	"github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg"
	"testing"
	//"strings"
	//"testing"
)

//func TestLitecoinWallet_CurrentAddress(t *testing.T) {
//	w, seed, err := createWalletAndSeed()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	for i := 0; i < 10; i++ {
//		addr := w.CurrentAddress(wallet.EXTERNAL)
//		if strings.HasPrefix(strings.ToLower(addr.String()), "ltc1") {
//			t.Errorf("Address %s hash ltc1 prefix: seed %x", addr, seed)
//		}
//		if err := w.db.Keys().MarkKeyAsUsed(addr.ScriptAddress()); err != nil {
//			t.Fatal(err)
//		}
//	}
//}
//
//func TestLitecoinWallet_NewAddress(t *testing.T) {
//	w, seed, err := createWalletAndSeed()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	for i := 0; i < 10; i++ {
//		addr := w.NewAddress(wallet.EXTERNAL)
//		if strings.HasPrefix(strings.ToLower(addr.String()), "ltc1") {
//			t.Errorf("Address %s hash ltc1 prefix: %x", addr, seed)
//		}
//	}
//}

func TestLightningWallet_CreateWallet(t *testing.T) {
	_, _, err, wal := createWalletAndSeed()
	if err != nil {
		t.Error(wal, err)
	}

}

func createWalletAndSeed() (*LightningWallet, []byte, error, Account) {
	ds := datastore.NewMockMultiwalletDatastore()
	db, err := ds.GetDatastoreForWallet(wallet.Lightning)
	if err != nil {
		return nil, nil, err, Account{}
	}

	seed := make([]byte, 32)
	if _, err := rand.Read(seed); err != nil {
		return nil, nil, err, Account{}
	}

	//masterPrivKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	//if err != nil {
	//	return nil, nil, err
	//}
	//km, err := keys.NewKeyManager(db.Keys(), &chaincfg.MainNetParams, masterPrivKey, wallet.Lightning, lightningAddress)
	//if err != nil {
	//	return nil, nil, err
	//}

	wallet := &LightningWallet{
		db:     db,
		//km:     km,
		params: &chaincfg.MainNetParams,
		defaultBaseUri: "https://lndhub.herokuapp.com",
	}

	account, err := wallet.CreateAccount()
	if err != nil {
		return nil, nil, err, account
	}

	return wallet, nil, nil, account
}
