package keystore

import (
	"fmt"
	"io/ioutil"
	"os"
)

var sampleKeys = []string{
	// 3ofcoins.net provisional key
	`-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v2

mQENBFVnXGIBCADovl1Yw/Ftjz7QZ9qVEUmJ9ztduws50yk+RbY+R1UJQvLhWU10
izORyN1qrXTYrQKoiJDJ9eTc7WHqrJGsNrgZWGc5SvMTdfVtiRdfaKP5QrIgXI5/
EZwZjRgi3ty/hsq2RvpNkvC5cXp4sSYf63dEFFLD4Ps3G+Lc4adKyNn8gZSieHjv
0+aizg7DW+mqNxUy3NK5wkRo876EYbhVpGCZRxPq39p88wqb6j1Rkt68jNtMOWkA
euoCxfeXR9jbU7ArNFxmhg6/ND7Zq60PAC5aImBLolXXMAK62cBPJBTZulNXSNxF
iZZBtFl5COyQNHSwbZUOv4jdrOiwE894qC8NABEBAAG0MkFDSSAzb2Zjb2lucy5u
ZXQgKHByb3Zpc2lvbmFsKSA8aW5mb0Azb2Zjb2lucy5uZXQ+iQE3BBMBCAAhBQJV
Z1xiAhsDBQsJCAcCBhUICQoLAgQWAgMBAh4BAheAAAoJEFPMwtY6FiZkYXEIAJiM
bGvBOIrjIKhW0y6syRIGuypkJw7dhp9EvMwzsvMLM86ze7iX2gNaBOGKWMdIRAYH
gKLN+f65ZOviOY5yy6fMG5A9ySlSWIKiLPDX6X/pXwMhmSUBkwhWKUWXIHd0+3d4
lGUWzzimTFfjWNlPWbHb08HMepMuNtCxtKkxvRgpH0yaWNCHQe3NsWAh6nYLsXpJ
A/2lomQ9K/qhHyZT18eYD1lHCpEOPksqifv5BKAFpXekGsFjTbYIKj6bgspaMg+d
GHMg3KUnCulgzhhbo0bMyMqrHm8waSF/JabTiIqp/ePFiUWqeirBdZ8yje6TA9eS
7UMOx+veKUT6Y5axbPaJARsEEAEIAAYFAlVnXLYACgkQRjqg8fR27d7nsAf3cxHv
Q+ljIyTqzxHDCgMmMLFPd8zzZ6cBTUYgb8R2G8EjYVg5Qu9zwv3FOZWaY19hBYDP
AzXjQtjkdQiwGvOppgui4CmnjJpNlyPd5J2A06fyeg2F28WJacrxUdTre1/UwN+c
9UUJo9o3OluVbaq1o9rXJnau79HOqWDoh8VdTzZAd3VO4SXuI7GUP49j6NC1t8FO
I7Qe0TLHzkZDeiZmVlK5OhfwhzMBQ5Qjv/9iT47Od91ryXKrHNAICyT32K4zawc0
/xs0oAQb5MOE4HrxW46e8UCT+2u+GrtS5ftWQqAMqnBF9oiRpSfXis3XgFyP5nU6
+ZaUESTObMoZyZ67uQENBFVnXGIBCADE19yLRCcO0NNBvDxWVFdFLsDKzl7vAQWm
cjaWBPJi0pWbi+vppECSjf8S5NNUXqx8MObG4i+ivBR4dqPbVls7Lpzn9wt3MwPZ
RbCLPLvBPJtlYtIt76JlqcjuJws3kd5hmX/1gny8Vy9RoNWt4y1ldde4+mJXGHoW
OY3CHj1vagkRE445czYT+ST8xeSxLKAyAEHis3U6Gb80uHbRknt/bUaogbibrl9N
3MyRX56cctVqpY8bJTi1PDKpSNcUHT+LMQ6rpPc2anKVqoUHT9Ye6hLSvNeBwJHs
LLWEWNRfmjjZvHs4hfzzaUVT0bP2IMkMBpxMCAOEargaJYhD1olFABEBAAGJAR8E
GAEIAAkFAlVnXGICGwwACgkQU8zC1joWJmTROAgA4RBWZww5CaBZ2ZbBijNRK8sa
swUrdxxcAkRzki2Z5KpBXESZEe6iPZt/rsGNw0F2XaORP+DRrzTGg+DSQlh5nT2z
LFVkNyeW3RJQjP2K2avikjRgoprryr/WVUOTpk4jK/MpBxUU0Htpvks2Jly2vYJD
Z9R2yW31RLEJXcVxOR/LcW6UfkjhOsH/RbKbcMx+HraxKk7jmrKZEglz8cO0BMGc
gemHpMb04W4YXOdKjqppm0kS+TDTEm85xnfAe14olr3ZsHn/ey8nc3JEb3aCBeuC
PjkMkIAW7BYX1ZKyV9z5EQVkqkz7zwyZYO1/tmdjQQH/x+mD0MwpeIF0E78YhA==
=m9io
-----END PGP PUBLIC KEY BLOCK-----
`,

	// quay.io signing key
	`-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v2

mQENBFTT6doBCACkVncI+t4HASQdnByRlXCYkwjsPqGOlgTCgenop5I6vgTqFWhQ
PMNhtSaFdFECMt2WKQT4QGVbfVOmIH9CLV+Muqvk4iJIAn3Nh3qp/kfMhwjGaS6m
fWN2ARFCq4RIs9tboCNQOouaD5C26/FsQtIsoqyYcdX+YFaU1a+R1kp0fc2CABDI
k6Iq8oEJO+FOYvqQYIJNfd3c0NHICilMu2jO3yIsw80qzWoFAAblyb0zVq/hudWB
4vdVzPmJe1f4Ymk8l1R413bN65LcbCiOax3hmFWovJoxlkL7WoGTTMfaeb2QmaPL
qcu4Q94v1KG87gyxbkIo5uZdvMLdswQI7yQ7ABEBAAG0RFF1YXkuaW8gQUNJIENv
bnZlcnRlciAoQUNJIGNvbnZlcnNpb24gc2lnbmluZyBrZXkpIDxzdXBwb3J0QHF1
YXkuaW8+iQE5BBMBAgAjBQJU0+naAhsDBwsJCAcDAgEGFQgCCQoLBBYCAwECHgEC
F4AACgkQcqv19nmdM7zKzggAjGFqy7Hcx6TCFXn53/inl5iyKrTu8cuF4K547XuZ
12Dt8b6PgJ+b3z6UnMMTd0wXKGcfOmNeQ2R71xmVnviuo7xB5ZkZIBxHI4M/5uhK
I6GZKr84WJS2ec7ssH2ofFQ5u1l+es9jUwW0KbAoNmES0IcdDy28xfmJpkfOn3oI
P2Bzz4rGlIqJXEjq28Wk+qQu64kJRKYuPNXqiHncPDm+i5jMXUUN1D+pkDukp26x
oLbpol42/jIcM3fe2AFZnflittBCHYLIHjJ51NlpSHJZmf2pQZbdyeKElN2SCNe7
nDcol24zYIC+SX0K23w/LrLzlff4mzbO99ePt1bB9zAiVA==
=SBoV
-----END PGP PUBLIC KEY BLOCK-----
`,

	// coreos.com/etcd
	`-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1

mQINBFTCnMQBEAC/49bGbStCpa3peej+/42mobfuGbTcdmcGGwYZmigP0Kl0TPZK
zcIxhVZOr3ITYuAx8T1WWJVb7/r/y4roUogDrUSTm2nAbLP8xp8Qn/N1zaXFyEtJ
WTaLuPI2mq643x8g7fAiJY19JRjFbYVVZLr5OMOvHrOdtYVN31ARZeSxmqP5yFNW
9DgoXG0/w80EOXIsWWoJgjaKLye15LHI86MjPthXyT5K212efxPffSk1hca04Dmk
I5vCMHC1Q2DbrlilhS0DTf+lSK2YgkaHPWiNSZb3XvjwbU8qQMzyfnQcQrQlfm4/
1fHLj2bWyASzNG/MOJCQ2JyEyIzbS2M4jKcfFxaKKuJA5PwdfjbRTkvxAKbFcdc5
ER7D3QoEOxgRDMppHaihKNI/T4dPIuqyUczq3ia9fGfrQFROAIxdAnBqzzBVaetv
FYFVjJlAhGsWWEhuO3P7qGwwR7CtPkWvvsMT8CYdHP2h7uOrOZioGCQ+09YBGpJ9
LzwCKHiV2s4/aBVfLhjttGqXG+PW/Kzg5rtwAjSdeooThKQLQk/ok1TtFydgNVNH
kSPNdhgiTWlNmK8Qj3C1zqZmcPzv+c6y6f79GTL0+Hz6gqnMGdIpFJtmbusU79/2
MkDqumBAslvwm7h85s0ccKwZCG1VelyhGLawVyxin0UhLWlzYd6SL5IuuQARAQAB
tCdDb3JlT1MgQUNJIEJ1aWxkZXIgPHJlbGVhc2VAY29yZW9zLmNvbT6JAjgEEwEC
ACIFAlTCnMQCGwMGCwkIBwMCBhUIAgkKCwQWAgMBAh4BAheAAAoJEFIQvYiIGCGQ
nxYP+wcVClXD1t5oT7mwvYZPfF9/+itOSHvN6++T4CCFkRVdIN7/G/Ou1wRb/fV+
P27Rc7gnK+jbQJqUa8aEsNSWZT/1A8VaQ51orQdV80ZROzrJPLBB0w4fkEsSESO+
Uuz9ZsiEOhZf0ATkafrF1jfepGXmoHJxLJ+bKS+KlAEhtceB1jiWAafGPy99XUAf
MBJ905F6bXvDqQov+9U3vyUGnwA6ymCCqKIoCVx3GdKOh5UaqC8pHaF7oI7xkv5b
5WMajzuXwKBy6KoduHTnW7Y5g1aoGwW5FDFoEq4LsBvcxeUI6OMVEWCVe502E+4S
lWx2gEvFa33wy77kYp2ZvHToY5tSjiIi8QocR0IgfLqE3P1ZMPe9YXk5EveEZH6Q
VtQ8z1ktuZCVQqrtTEeernEdSFsTVFSoWUsNJV1FMlgisZZ0ljoFJrH7/7GkYhK9
DT7OcrZnyZDUkEIiVaqWwjWw5Ing4IHExAq+PlXDwrA0QcH2UW9IurDCvPiXHpxi
D0V21oEbANUdGtYcOSDRBbZlsonINJZQ1Ad2XrY8kfZ2sZXAZuBZbH0+dEx6zmua
C0IGN3SLFseScqJZ5G1joYYqOKOUweErkzA/62Kaj31SVoQDpZyMqtwTjjZaFT8N
fMkBtaM3knaFonHZc19BD1FOishRThCCq2Ty8HUoN2Fk7w0l
=bYl7
-----END PGP PUBLIC KEY BLOCK-----`,
}

var sampleKeyFingerprints = []string{
	"4706dc5d5c214bc3ad127c6d53ccc2d63a162664",
	"bff313cdaa560b16a8987b8f72abf5f6799d33bc",
	"8b86de38890ddb7291867b025210bd8888182190",
}

func asFile(args ...interface{}) (_ *os.File, erv error) {
	if f, err := ioutil.TempFile("", "jetpack.test.keystore."); err != nil {
		return nil, err
	} else {
		defer func() {
			if erv != nil {
				f.Close()
			}
		}()

		if err := os.Remove(f.Name()); err != nil {
			return nil, err
		}
		if _, err := fmt.Fprint(f, args...); err != nil {
			return nil, err
		}
		if _, err := f.Seek(0, 0); err != nil {
			return nil, err
		}
		return f, nil
	}
}

var openedSampleKeys = make(map[int]*os.File)

func openSampleKey(i int) *os.File {
	if f := openedSampleKeys[i]; f != nil {
		if _, err := f.Seek(0, 0); err != nil {
			panic(err)
		}
		return f
	}
	if f, err := asFile(sampleKeys[i]); err != nil {
		panic(err)
	} else {
		openedSampleKeys[i] = f
		return f
	}
}

func closeSampleKeys() {
	for i, f := range openedSampleKeys {
		f.Close()
		delete(openedSampleKeys, i)
	}
}
