package threefish

import (
	"testing"
	"strings"
	"fmt"
)

func TestThreefishInternal512(t *testing.T) {
	if !debugThreefish {
		fmt.Println("WARNING: Skipping internal threefish tests (debugThreefish = false)")
		return
	}

	tests := []struct{
		tweak Tweak
		key   []uint64
		input []uint64
		want  string
	}{{
		tweak: Tweak{0, 0},
		key:   []uint64{0, 0, 0, 0, 0, 0, 0, 0},
		input: []uint64{0, 0, 0, 0, 0, 0, 0, 0},
		want: `
:Threefish-512:  encryption + plaintext feedforward (round-by-round):
  Tweak:
     00000000.00000000  00000000.00000000 
  Key words:
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
  Tweak schedule:
     00000000.00000000  00000000.00000000  00000000.00000000 
  Key   schedule:
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     1BD11BDA.A9FC1A22 
  Input block (words):
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after initial key injection]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after round  1]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after round  2]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after round  3]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after round  4]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 

:Threefish-512:  [state after key injection #01]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  00000000.00000000  1BD11BDA.A9FC1A23 

:Threefish-512:  [state after round  5]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     00000000.00000000  00000000.00000000  1BD11BDA.A9FC1A23  C178E7C0.8AE7CB38 
`,/*

:Threefish-512:  [state after round  6]=
     00000000.00000000  00000000.00000000  00000000.00000000  00000000.00000000 
     C178E7C0.8AE7CB38  1BD11BDA.A9FC1A23  1BD11BDA.A9FC1A23  ED9BE223.15E5E0A7 

:Threefish-512:  [state after round  7]=
     1BD11BDA.A9FC1A23  C178E7C0.8AE7CB38  ED9BE223.15E5E0A7  1BD11BDA.A9FC1A23 
     C178E7C0.8AE7CB38  E5DC0A57.4171F777  1BD11BDA.A9FC1A23  C29EDD4F.CAF4F808 

:Threefish-512:  [state after round  8]=
     DE6FF92A.74F1122B  A5ADC311.D328DD9A  D377EC7A.5757D81E  2BE07C9D.BC25111D 
     DD4A039B.34E3E55B  A4923070.0016A9E9  DD4A039B.34E3E55B  89C8396C.6007F855 
`,

:Threefish-512:  [state after key injection #02]=
     DE6FF92A.74F1122B  A5ADC311.D328DD9A  D377EC7A.5757D81E  2BE07C9D.BC25111D 
     DD4A039B.34E3E55B  A4923070.0016A9E9  F91B1F75.DEDFFF7D  89C8396C.6007F857 

:Threefish-512:  [state after round  9]=
     841DBC3C.4819EFC5  B37B1557.38DD9B0F  FF586918.137CE93B  3D0978CA.AD7B20E0 
     81DC340B.34FA8F44  025C34BE.7BB7ABD5  82E358E2.3EE7F7D4  821C5213.07E0DA58 
`,
*/
	}, {
		tweak: Tweak{0x0706050403020100, 0x0F0E0D0C0B0A0908},
		key:   []uint64{
			0x1716151413121110, 0x1F1E1D1C1B1A1918, 0x2726252423222120, 0x2F2E2D2C2B2A2928,
			0x3736353433323130, 0x3F3E3D3C3B3A3938, 0x4746454443424140, 0x4F4E4D4C4B4A4948,
		},
		input: []uint64{
			0xF8F9FAFBFCFDFEFF, 0xF0F1F2F3F4F5F6F7, 0xE8E9EAEBECEDEEEF, 0xE0E1E2E3E4E5E6E7,
			0xD8D9DADBDCDDDEDF, 0xD0D1D2D3D4D5D6D7, 0xC8C9CACBCCCDCECF, 0xC0C1C2C3C4C5C6C7,
		},
		want: `
:Threefish-512:  encryption + plaintext feedforward (round-by-round):
  Tweak:
     07060504.03020100  0F0E0D0C.0B0A0908 
  Key words:
     17161514.13121110  1F1E1D1C.1B1A1918  27262524.23222120  2F2E2D2C.2B2A2928 
     37363534.33323130  3F3E3D3C.3B3A3938  47464544.43424140  4F4E4D4C.4B4A4948 
  Tweak schedule:
     07060504.03020100  0F0E0D0C.0B0A0908  08080808.08080808 
  Key   schedule:
     17161514.13121110  1F1E1D1C.1B1A1918  27262524.23222120  2F2E2D2C.2B2A2928 
     37363534.33323130  3F3E3D3C.3B3A3938  47464544.43424140  4F4E4D4C.4B4A4948 
     1BD11BDA.A9FC1A22 
  Input block (words):
     F8F9FAFB.FCFDFEFF  F0F1F2F3.F4F5F6F7  E8E9EAEB.ECEDEEEF  E0E1E2E3.E4E5E6E7 
     D8D9DADB.DCDDDEDF  D0D1D2D3.D4D5D6D7  C8C9CACB.CCCDCECF  C0C1C2C3.C4C5C6C7 

:Threefish-512:  [state after initial key injection]=
     10101010.1010100F  10101010.1010100F  10101010.1010100F  10101010.1010100F 
     10101010.1010100F  17161514.1312110F  1F1E1D1C.1B1A1917  10101010.1010100F 

:Threefish-512:  [state after round  1]=
     20202020.2020201E  2423E424.2424241A  20202020.2020201E  212120D1.2121211F 
     27262524.2322211E  8F86BDB4.AB5A99AE  2F2E2D2C.2B2A2926  2D2C2CCE.29282B24 

:Threefish-512:  [state after round  2]=
     414140F1.4141413D  0C0C4C70.0C038C70  44440444.44444438  C5C53C75.C5C205B9 
     545251F2.4C4A4C42  11D9C036.70EF6135  BEB4EAE0.D684C2D4  251B10AB.6D232D24 

:Threefish-512:  [state after round  3]=
     531B0127.B230A272  F8BE8665.40ADC0AA  695F14EF.B167715C  8F09ACDC.E4AD4309 
     605E9E62.584DD8B2  5DED1276.AFACA115  847A2756.9C46C88D  F8C986FD.3CEF24EA 

:Threefish-512:  [state after round  4]=
     4BE48824.EF1FC75C  A1320230.3492DD3D  C74C2766.61141271  E6E74293.E01FB6F8 
     EF684B3F.3CFB1BBB  821B5C22.FCBFF959  7D38ADBB.DCF48937  D8E9725D.315612AD 

:Threefish-512:  [state after key injection #01]=
     6B02A541.0A39E074  C8582754.57B4FE5D  F67A5492.8C3E3B99  1E1D77C8.1351E828 
     2EA6887B.783554F3  D86FAE73.4B0C43A1  D48F0310.3046DA87  F4BA8E37.DB522CD0 

:Threefish-512:  [state after round  5]=
     335ACC95.61EEDED1  E925E271.4DFD74FA  1497CC5A.9F9023C1  1043B650.98177E33 
     071636EE.C3419894  2B273869.A2FF2159  C9499148.0B990757  FE92C364.DB6DBDD9 
`,/*

:Threefish-512:  [state after round  6]=
     439E82E5.FA065D04  41F38774.4312C59F  FDBDAECB.ED8D98BB  2F3FB2CB.06607D83 
     05A8FA53.9EAF566D  68916F3A.521D4C1C  F470C9B1.AE9828B0  F2CF0018.933C3BDB 

:Threefish-512:  [state after round  7]=
     AC2FF220.4C23A920  AF1AA44C.DF41FB02  F08CAEE4.80C9D496  437C73CC.D11F506A 
     479C81C7.E1C21C0C  A289FC14.04943409  23B07C7C.B4F8A633  11527172.F8C9100F 

:Threefish-512:  [state after round  8]=
     BD826393.44ECB92F  C86F6C16.D5C1A39A  9316AAF8.855E089F  782CB240.A871B36A 
     8B18F594.B2E16C76  9AB42304.915A9CAB  D2CB20C9.943AA135  7BCAE3EB.CE7F32B8 

:Threefish-512:  [state after key injection #02]=
     E4A888B7.680EDA4F  F79D9943.00EBCCC2  CA4CE02C.B89039CF  B76AEF7C.E3ABECA2 
     D25F3AD8.F623ADB6  F20A7858.E4ACEDFB  F5A241A8.4138BC57  92E0F8FF.E19143CA 

:Threefish-512:  [state after round  9]=
     DC4621FA.68FAA711  2F769C1D.0EAA672B  81B7CFA9.9C3C2671  BB090582.EA92D1BF 
     C469B331.DAD09BB1  06AE9656.B50F0BE2  88833AA8.22CA0021  BAAB43FA.7ED51FDD 
`,
*/
	}}

	for idx, test := range tests {
		got := collect(func(){
			encrypt512(test.tweak, test.key, test.input)
		})

		if got, want := strings.TrimSpace(got.String()), strings.TrimSpace(test.want); got != want {
			t.Errorf("%d. got|eq|want:\n%s", idx, sideBySide(got, want))
		}
	}
}

func BenchmarkEncrypt512(b *testing.B) {
	var (
		tweak = Tweak{1, 3}
		key   = []uint64{0, 1, 2, 3, 4, 5, 6, 7}
		input = []uint64{0, 1, 2, 3, 4, 5, 6, 7}
	)

	b.SetBytes(int64(len(input))*8)
	collect(func(){
		for i := 0; i < b.N; i++ {
			encrypt512(tweak, key, input)
		}
	})
}
