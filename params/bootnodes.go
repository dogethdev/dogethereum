// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://de1a8cf1f903ea3cf17f15927d9f0bface23fc66eb20d86c34364c5d335e54830f72a5b73742212ffbda37ab0ccb27627cd19b71705c6be92548bd09a4e02429@35.236.80.146:30303",
	"enode://4d80c47e075de00ad3dfa0525568840e24b2d5c1eae996d3582bed2c3cda26abdc5f934e139a3e8e00af9fad434fa15f930623f7ce7f4056d7e962fb9752b793@109.108.71.107:60528",
	"enode://043aa5270315b8a2353e2fc9815bfb62233398b2e21d9a6d8594ee25203026b1e3ca710a0ef4d03969ace178686d033b4876e9bc35965638bdc0a77492498fdd@54.37.136.64:60084",
	"enode://9dd6d9427f65eb730c6b685b39cf548b05a98c0c3ab48944da6c912e29bc6b4896f3e3ff2b85d5ec87c3b612206c9686ef4158fe514db4cf0795a9e0d615bd73@94.198.55.209:30304",
	"enode://9212134279c09f989c14dd4320581329677b1a090a8978bab99aafb11d058521055632c5bbbe3e88c75e79b100afb420eb4dcd7499d3a5ae6e645816e1058b1b@89.108.117.245:38692",
	"enode://cd63f6b8a33dc3e9abe87c6b436560df35fe2d856ef09da6b59393bff23c97eb9dfc0ddd235c67824af7f183f6cd9a89f04cbd9e22ac0ed135a4dfa9c64ec2f0@8.217.59.184:30303",
	"enode://8c39bec0260e8be213e4f4060654a5dd5c12570e031d56aeeeb8f70d3518b7ade662ad2453c38192a90ebecd50f54e2d908a6759cb5ae19231a6a73cc9d999d9@15.235.212.164:53512",
	"enode://3d3d287e3b2697236ea7e5c8d68f756550269af58f21612d2f564e73b366a9a622a1fc75d2539f1a63445bb70e52426a851dbfcec5c60f35d1e21e23a7a390f2@188.163.19.44:1222",
	"enode://2d8864f3177c9a6b2f02879e9b8726ef5ab401b049af9e330db490a50367bc294910ab773bb8e76f51936fc53f0b62a2a4dafc4c13f0ffc5fdbadc655839134d@15.235.182.112:25301",
	"enode://6100df7b3e4601943b5f0d6ea42be766b01abd41945b28cf43905ba383dae3da8b3569be7edc51b426f523ed0b46b267e418d5043aa44bfec73ccf0d52014622@185.183.35.91:36058",
	"enode://f55f842231eb43f336bf2e7cec0f5ba6ca24390e35a32ab066b1b4cf51d2416356894394bce2a643c83ab2b631b26217365006d437cf329cf43e11088d343c0e@82.74.183.87:30303",
	"enode://135dbf51a945df9334e3787f4cf9ae27ce830735b768a8a4b0547034a5e183cee3a6b49e452a7d33c6f2634af371e6feae4d7d690f3376dd33d3293ce55cacaf@5.9.151.50:40302",
	"enode://b2a27b2f4d72fa33f8b467047ad906eb3e91c00545bec840d68dc24c3bb6e97f6d325393e4876d33dbf94f3660bedfbb66a1e6be6758829d974177edb4a5712c@89.232.109.242:48938",
	"enode://7de408607930d8fa06e7ca2c51fa08acf3ac68c0e0c6f78abc16b84a4cb93d0374c880db17452a9e36dea70b3276dd96df0bb5af88b48ba8477f0ba5e0355ea3@46.4.41.220:35246",
	"enode://32961d601f8978fdb1358478f969f2b8663b43c50292612e85b47af9bacda9c3c6e52c808980e9443c96a9744f432cd9ec4691eef29ee5bd3d21ac0d18d37844@91.185.84.236:30303",
	"enode://3219d99ffc191aa22ad8c009224407c49f6a2a2b7591ca64e60dfb9f7650da878e6185d64e166ae156318cfae729180065b8eda4e0260d601894f4fd22fbd9cb@148.113.0.116:43842",
	"enode://00084be5030bd13d5cddddf01003a5a883cac028d67e22a2b87f09b60f7d3f96f671e128e1d32e5927ca6bb57041ccdc3f88067c4f48d626cfd352d8e8be602c@194.180.19.36:30006",
	"enode://8162098ea6116f76bee89c0caad1cf5daa09c6a545e658947f5d6f79b7dc7cfde33c46a7c0bbdfa3d4143c6644482cf3625b10e6d2b4f023b7f84e6af0363396@89.232.109.242:41604",
	"enode://185b1b28f99d0581160716fec6a11394b980fc4d8cee363982d0444856676e6e2322f5d43c944595532e427b200b3392faecc4cbb5ac213453095daca7bbf801@57.129.28.9:2298",
	"enode://c03858e6ce410b356dc67c8651dc1e31b721540dee2aaf265e926e9e42775ba4503ac77d14b16c121729e890e50bd08c2d8904ca6580a41aca8f4bff2cf471e3@147.124.222.176:52706",
	"enode://08259af41d4c91ba756934989a2e197f8dc0cbd9410e998d6e2b278f3aab1c8e364065341b14b4559a2c1d6bdc268892cc17446cd1c7ece87cae5bcfef21cbf8@116.202.50.249:30304",
	"enode://615c033e5b72e00480e9b449f11ed5ab9c277064b7fae7f61e9f06a0b56be061d1fbd2ed33591643ce2c7e24face07c45d49110fb88c622ccb5605462fcd03e1@104.243.40.223:5443",
	"enode://447de79bf5d49419b2708a6b40671657c5756ca8ee1471ebed2be12a000f2e72ad007d1ce0739586558b68208dea8878352cd654d32a2fa3f7a6d55a450ec017@15.235.55.110:30303",
	"enode://8108bf49bdf7941eeb8372f460e875fb5bfa136a1f0854bc6c8adef1aeb7be3996c6f3c3890340b237398241237180a9b83ecd2387442795d0b2095f10db82c0@5.188.238.193:46868",
	"enode://960bf86c27a73965cca275be90b1c1f13266a65875db675a66d31ad6b399c10896ddabbd9a4582292a7fd0bc822413312cd9e304b3659fc3132715908f807cd5@51.161.131.90:58654",
	"enode://aaf1ad0bd83aceaef066fa636115392d34fe73a5d670edee84fdef17e28f5c3ca2421f9699e5f2cf57c07bec7531278a940a8435e8e98145210b7774024943a4@135.125.8.172:64086",
	"enode://a1fe274aee9c27b2549bacf6623233f22fffce521773539fdef5849459d1ba094f5a1da327a6e1f193f65d243d01856dbb94f86e05d72026a3239cd59fb5ec66@89.108.117.245:49770",
	"enode://cb4014afba9dfc4890568b28f91ceb8e40a0e36b63b73d3dc7c4791eca28f8b7aa73960076801b3f4ad77794e78f80ee78c2902a17ece0d2a6ba2be1ffd477c7@65.109.52.145:34230",
	"enode://e6f735f33c769ad700a3e5aca23a4999bae97cf60077555b597bf788ae32ad62bc663a21b1f2223093252f8d55a0ce5edd736a4c4556eacec8ca2c8395bf6f15@202.181.200.173:30303",
	"enode://d96f7556fe85c16be19d4542414cb5ebc3e1fe14cdba7aa55a5421f3f51072756f4564a1447978bd02c1f4c8a66a39c4e26374550ff1d9d7e4b7e9258e6e7386@35.236.9.214:35678",
	"enode://dbf9cd1d0f25393e5217b59c901f14ccfd4d5423139e286bbc765f5afb8f13125dd7fb2de95808c5b0f9523c15704f5f82df6dda53b4e4510933a16a35a43b94@185.136.76.96:39272",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// geth
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	// besu
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://a869b02cec167211fb4815a82941db2e7ed2936fd90e78619c53eb17753fcf0207463e3419c264e2a1dd8786de0df7e68cf99571ab8aeb7c4e51367ef186b1dd@51.15.116.226:30303",
	"enode://807b37ee4816ecf407e9112224494b74dd5933625f655962d892f2f0f02d7fbbb3e2a94cf87a96609526f30c998fd71e93e2f53015c558ffc8b03eceaf30ee33@51.15.119.157:30303",
	"enode://a59e33ccd2b3e52d578f1fbd70c6f9babda2650f0760d6ff3b37742fdcdfdb3defba5d56d315b40c46b70198c7621e63ffa3f987389c7118634b0fefbbdfa7fd@51.15.119.157:40303",
}

var KilnBootnodes = []string{
	"enode://c354db99124f0faf677ff0e75c3cbbd568b2febc186af664e0c51ac435609badedc67a18a63adb64dacc1780a28dcefebfc29b83fd1a3f4aa3c0eb161364cf94@164.92.130.5:30303",
	"enode://d41af1662434cad0a88fe3c7c92375ec5719f4516ab6d8cb9695e0e2e815382c767038e72c224e04040885157da47422f756c040a9072676c6e35c5b1a383cce@138.68.66.103:30303",
	"enode://91a745c3fb069f6b99cad10b75c463d527711b106b622756e9ef9f12d2631b6cb885f831d1c8731b9bc7177cae5e1ea1f1be087f86d7d30b590a91f22bc041b0@165.232.180.230:30303",
	"enode://b74bd2e8a9f0c53f0c93bcce80818f2f19439fd807af5c7fbc3efb10130c6ee08be8f3aaec7dc0a057ad7b2a809c8f34dc62431e9b6954b07a6548cc59867884@164.92.140.200:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
