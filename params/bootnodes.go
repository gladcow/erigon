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

import "github.com/ledgerwatch/erigon/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",   // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",     // bootnode-aws-us-east-1-001
	"enode://ca6de62fce278f96aea6ec5a2daadb877e51651247cb96ee310a318def462913b653963c155a0ef6c7d50048bba6e6cea881130857413d9f50a621546b590758@34.255.23.113:30303",   // bootnode-aws-eu-west-1-001
	"enode://279944d8dcd428dffaa7436f25ca0ca43ae19e7bcf94a8fb7d1641651f92d121e972ac2e8f381414b80cc8e5555811c2ec6e1a99bb009b3f53c4c69923e11bd8@35.158.244.151:30303",  // bootnode-aws-eu-central-1-001
	"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",   // bootnode-azure-australiaeast-001
	"enode://103858bdb88756c71f15e9b5e09b56dc1be52f0a5021d46301dbbfb7e130029cc9d0d6f73f693bc29b665770fff7da4d34f3c6379fe12721b5d7a0bcb5ca1fc1@191.234.162.198:30303", // bootnode-azure-brazilsouth-001
	"enode://715171f50508aba88aecd1250af392a45a330af91d7b90701c436b618c86aaa1589c9184561907bebbb56439b8f8787bc01f49a7c77276c58c1b09822d75e8e8@52.231.165.108:30303",  // bootnode-azure-koreasouth-001
	"enode://5d6d7cd20d6da4bb83a1d28cadb5d409b64edf314c0335df658c1a54e32c7c4a7ab7823d57c39b6a757556e68ff1df17c748b698544a55cb488b52479a92b60f@104.42.217.25:30303",   // bootnode-azure-westus-001
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	"enode://7c9740e4d64674801fe62b76798d46778a038c49caebb15843d8c0f2b2f80d7ceba2585b4be366e6161988f81ddcfcd6fca98b5da52ae9a6f22c1b2a84b24a04@18.130.169.73:30303",
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
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

var BscBootnodes = []string{
	"enode://e8445158a924fbadabc5cc6bf0ef4059d544279460eea5674c586db24bf93b2452d5d911f9db67b3b4205c04908b64175461e6f64207729d661b401d3d0e7849@3.84.218.95:30303",
	"enode://03fca5e5684c2bb3302a7fc56d979c2ccb1b44f6a755eee12cb40ecf40602c4aa71e001e74bf83905403977ba67c2d4ddebdd8ef4b20f3d1777974e673e65db9@18.181.197.121:55420",
	"enode://e18b67f01b19555b03930c82ea07424399e981daf7fa9abfb00a94a144d2a4e3e1468978815d41ba0755ed54431215cf3496f4a3ae118a024f956cf66f21da59@3.123.156.202:47008",
	"enode://4bf1ca233a5ae05e889ace6723bc53b0815b0288f9cb06338c1d3734e37b7812ae00a87e1b53e60cea307831f372615515f18e6f9b24edfda89aeaa0ba9566cb@23.20.242.79:34042",
	"enode://f721ceed876d4358a7b5c882717dd17230ae237db16b2215bccec18cfcea62e63027aae80407b75fe727a1ff1caae4372f417b90e920f3d3d57c903e34118993@54.86.168.183:33978",
	"enode://88c2d74e121b60fa171881561570aabb2dda4c18f620ff93c8127871ef59b154e0025e978d103e98b4062474c3897111e95fc8defc7a0fb7e1b5dd191fc19eab@35.153.242.98:31810",
	"enode://ac182ff00b369d944644d6498ac9050a2b6945d152891df8afe3988d425c8617fa3edef461539249d182c40dfc4e33f5130354a65064b511c749a66997ac3bb6@52.22.25.99:44696",
	"enode://042941f93619bbd9c30ffa7e55c1d9bb4bbe8dfb3ed386a8d10bc013fe530203c659ee609dee7ec5a1469b6528522cabe8f24de145b6a3d6b55d441fa5dab794@52.73.145.66:35080",
	"enode://1ec9c098f2234987c9b74cc8de13d9072685e59e8a764d2324222108f606361d201f5b0102c806833f715b1b2eb78c7b9485843b921b1ffedc196a42e02fc4e4@52.199.207.147:60748",
	"enode://3110eac4bb6c35c87bb27b6607b6177c5174cf2dbc8f1900fc7ff95e5e0d02fbe571e956399b3a2411ae1bb33d6fef09ec31861100c9bf413fb580e40cffe50b@13.230.42.250:36740",
	"enode://2e13ca7f602b1aa57a556d20236f6173e14b823ecc5c83794249ff726c1a33063c9e9866ab87c644c9cf438030f1dbea315f373e92b7956f0f5b547f2aeb28ac@3.68.67.182:57356",
	"enode://ed97c03808c5c2e85f88873ce9d7ea0e13f25ad2ea34e7166ac5ab394654566463e5b4e89c6fc82d698e4f723c15c6e8f0ba637a48aae1b5cb2fdf2bef7d6a68@66.135.1.68:30311",
	"enode://f3cd7b3c2f3fe8083c4b7f062e66cb048e32f4bd81c3abafda08c283d804121a90d3f27a680d10fd77d9d37c79affca8b3d24995b9be44a47849868f58bd8bfb@44.200.133.198:37920",
	"enode://ca078d6849de674fe7fa0a7ca55057978566499d2c7401739d8ee6a8933a3ac3e3c29cfc6f8474e86dd576035ba0d92038115917f928d43c86e01eb761cac912@63.33.196.130:53888",
	"enode://bde8a6bd27702b69480f9d4f75e7f89386b3252eba03f1b76df29595e483bbeb2770163bc109a114cfdc0cabfb775da9cd4fe8e3a09e39aa461ec4179798cd03@135.148.103.41:54842",
	"enode://00c32406ee438fe7bf3ef822b3419726c802d443e785189f7b276523c196592a35541dc12c45a730f62d41a92deb032eefc216517843f2f9c30df7f80450ab9b@65.108.47.120:57776",
	"enode://7d6e01f5f6a3620ece9b1cc66ef1a7e15dba20592c0bc45c22fd1960086990eaa179675e6c88c84ab90643b847aa007ff9929988626867120dba44c2a1df2dd1@3.231.93.45:40840",
	"enode://ca85e4425ebde5bae3ea261249c623d77e8d86b588ed98f20005972587be1f6c771b9b2b79e98f46b0e0cc25a8b8404d5cf31ebf1ece7a2eb8f19bfbf30e7808@52.195.1.108:47192",
	"enode://cce9c930ceaedc17602d42f5d567f4ce919c60f94c92df9a34220ddb68f566b239b58cb1ddc210b087dc13e35f9703716f8feb05720903920298568ec72ba729@65.108.97.34:42706",
	"enode://cac2caf1e6ae3c1c3d1368adf54744e2a5d6ac09f4510cdfeac29d67da7541d5419215d1eee3dc3463e5b4226adc51fcf055867e43ca590c4a70a9dca1ae0554@54.95.191.100:38378",
	"enode://d9e3dc8ffa71ed758cf593d8ea96f61960f63a17cdf9e0513836b848a28aef3491c005d67a8be3a0f7398de5da10369ba842b53d6b43aa9558d5fb08b0d6eeee@3.238.193.187:60102",
	"enode://97865befe5bf355f09961200108ed5bf7657d8463046f0555d09df3d4bf35e53a970ea23c6df6b77b2cca8fa2aa7030c65e73d460fc2bb933dcb64d7add42ff8@13.231.15.37:37518",
	"enode://f917d14be8314ef4b724b6217efdd8557580e159e3b488523556f8decbd00be6c7e65e22aea012bdc636be6a1bb8f4319c065408950e2285192ea6676e5a6a78@52.23.214.142:57920",
	"enode://b81b738a0f360d1f23bee81b7ad089ffde6caaf506bb0fc3007763f4b31d174744f9bccaa415c836d017de3d888c1c97836d4d79f76c2453fcbac85c1518f8da@34.159.158.66:60332",
	"enode://153807abab5c64498f0ac3ba4295269326091df860a27777fa32c6cfb99f71951b0c96f9313b3d7103cc61b80e9348fa56d9bf5c12ab65057c066d3a8ff88fdb@95.217.226.78:48870",
	"enode://eeec3cb8fbaf37e07c577ef40b86a51262e5702d5d9b90ddc2d9778b117f9495564909e8024c26a64e2181104bdd60907da665a32c7854832e7645fca7e59f70@65.108.126.143:37056",
	"enode://7513e884ef20bb2d04a90288cfca821d94045481a4fc79b30e45f9e4914697cc35d042060587311d12305044c54d8479eddff8dd60b23ab2d8c755e691aca480@3.93.239.226:40006",
	"enode://a123f1703cbfe431fe8a5f04abcea1399124976a901afc4fa090520f1cbb3b673998f66d212e1e94bdb5dd6f729d71721ea9c8b15fd4e2101276409bcdc23091@3.88.169.36:49352",
	"enode://b3b89948a6f14d9ece5258d45f430d3110e1934729a157ce7a021f8b9ca14d18289da0460d7e35998b742b445f624b3818da3019b67e4f6f971a9af81f9e82f9@34.230.181.216:49230",
	"enode://57336348367a2db45e6b420494d139423f48e7238fe0d2cf63fc6676d9da76e52ae9da4fc8f3bfa4a8716aa51907618bad5ab8dcb9eb7b9ab3d2999a672a0558@116.202.52.95:46386",
	"enode://d350de2adce4ec31a9df94920ff1207aac4f4c5486473f0a307b8e80bbd8ee896d32c9cc3ba6a3b8b25de147c5e94e1809e1b833f90cacbab0de7e7bd9abcb1a@54.167.123.85:45872",
	"enode://d54499eb86df0f28d91906cef4f648360b636f923d58567865b83508fb97d2e95d2ce3e3a594facb73230b9e19172a534ce5af3f528599d0d1e52632438aa864@44.200.226.84:32846",
	"enode://3a1107477d256c3490e459bd17c09187000c3899eca0e0c2d57811171015d213f0adf7de97d66b8b55f4df2cfcd28fde8f492405fe0c936fe86fa485cce00314@54.234.72.137:34502",
	"enode://ee3dd85152c8c208674fcca0d3e11816f9be632e1eacf321e9754374d30e162ed32323ec6abf4a0de7e9563a307d2a469932d42d7960fe3216ad237fa0416d7a@18.195.214.70:32976",
	"enode://089fae127dff1e7d418e65eb3c26a91d1090ea586a4a1a9d686daa02195341817db39cffdabe558eeb5e3cd355ec6d054746be6f3c8a131aba8605f9f1c49e79@54.73.165.132:53380",
	"enode://87725dde3f71300b6461f0444c060d04853619fb86abadaeec66ad85e671c839f8645f6bdf5e74f017f5d9ae67d2546bf82bf972db9cb58f1bce29f2dba1f76e@54.195.184.10:35598",
	"enode://2e37f5d37dd11ed7d9ae81ffb448574e331ea018f0065044e2f6ed44769b84288c30e138ac81ab98cc2cfa799e59c75d1412346112a50f5f540ba6e957e10fe7@52.23.162.149:53686",
	"enode://ec6840d6e2111ba8268521feaa4377b565459b80c3d4063a17661b77359dcbca33e454c8099b24665c5712312a2ef61a73dae28f83cca015e09b39e8d4108d8c@3.115.20.76:38292",
	"enode://5a3dee4b8db295745adb793b8046d99ac7b8dadcb04987b1b540801e9b2fb67a611de8ff684e29624e25e32d87e5a2fd8402526cab95f57311e69e1385253c8e@52.91.214.53:50594",
	"enode://430ae1f5473bae9605bf95ed8748549ce4010d610b8c89a42966f0fb9ea1c28016430ec4aa8ce1b87a51172f8c1944930e38224497e877546692e7983084a75c@65.108.47.119:49984",
	"enode://038d68b4c45f6dc361d9ab585c305f01311d4e13e0042e26dcdda7fdcb41ff631621db07a599e02c1e0a13f573b335b3b7b1b8eb98b271b139d8bc3d42785922@168.138.165.83:57346",
	"enode://9e2660bc70de1ba7de5349dad85d5811a802225773bc9e4be2184983b1c64af8c6223f0024cb856f4199fbee2a77304874eaa307f750d211fbcbf54307829f2c@107.22.106.3:53262",
	"enode://25e2d9e58ea341243780bf01586eda67f31968138f9646b669fedc2ab3c9f59ee908d9be39f06c7e8dd5e8caf412a2d49bc6d18ce69ebdbf32fd70494548e621@35.72.181.22:46954",
	"enode://9eba427d9d11240af112d11ee392a03d842903e84e36b8fb6629db5933a6639522f57e8093ad5d858640d57fff447c2d33ca1ede7ddb4f62b1f4d5dd41c0626e@3.112.247.160:43562",
	"enode://b02bead6b000cbea06aa7eb06302a7a5b6c77f9aee2ba764c73540198b5606dbe8e3cc771b092dd991e47aac6461fc88f0f2aec49f8a29065738c08ee4509071@119.8.25.79:43796",
	"enode://8bc72badeecf51e770ab707bb6523782c55fd163ed26068a553f1683769e3ec7364c616966b8a8b08fee6f988ce2da989be2ed5e1ce98b848eac3cc86136c3e5@3.113.247.226:51860",
	"enode://322a42a08959aefd3423d17d8aeb802e0dbfb8bb0096aa712b6bf3036c91a80b0abc45c7a3d1320eda9a9c0337dd028967e4b84357080c258c8d0a3aaa02a821@34.245.12.138:43810",
	"enode://e6fad1811d9e16d0b165f3c0e85a5d5141b20e78c1824ecf859bf52b1f834ed1454d0b7ba0abefc093ce0dc97842bb615243118d23363a008bfcda537be85f32@54.162.124.62:60590",
	"enode://178b2c581f118632465d2eef36bba16bf6bbf9d36ff91b4f7c95e46f8e550d264ca8f72dd5d605d89c32957d255122bcae574fc0fee96b7ab599b16ce073c730@44.192.121.160:59734",
	"enode://52a8b83e7e6ca49cd1fb488d55fe29d0b910f36eb2e225d7ae1741b9ea934df96775708cf516efda532c4e7f2ed60e8acdca3df77a57a21ad918a9595588985d@137.184.72.42:36698",
	"enode://94073ad5415c079dddd499f85b162f928858c75a457f8fea80d16e08f1df77409e914d25b47dae0ae8ae38a2ce03319ca646f3e0d6e687bed37ee0aa884da81c@3.235.67.204:43482",
	"enode://a223f411a466423a66d214bbfd385500857413c30dcf1ee80e7d1d551aad9a6d9113609fa32162adfc65e293b5cfb679993c4f16e80ae170e6135909eb0be422@13.40.7.37:48334",
	"enode://9f005be9111a6152884fd575abb55bddb1e7f726510c96cddde57a9bba84ffa4952a89d7632c9c9dd50d3750f83966a73a0f7ed793f253a3691b84a687b29b6c@3.88.177.211:35382",
	"enode://4e6170d4a0f1da531b4d48c003371e9f5e8e2b93fef9d9021d9bf6903c6566024b3bef96d6bc5b9b8150f6f8d02f29d26238a372d4d660e438de9525a733454c@18.156.192.65:36544",
	"enode://1f5a6648a97f69c9a4d2b3fe154b453b05746f79aeef4c358ef40c626dd7a814f2ad9a41f7db84cbf4ef21be08d95e625a7e0cd9e96592df375916c465a11046@148.251.232.16:54930",
	"enode://5451251a9902e658154456ea98ebdd93313e54496ce0a6ca2242fe4db882940d78d758c85a36485af54b0841270f2bdbff64d66c45976f3ed1dd912f7649c831@3.236.189.129:60464",
	"enode://df04a45098cf4745d504e4613cc5a9aa171fa3278bbc29c1387f7981ef27169d4290acd0b176fa66e3d435cc044074a4fe704b956c1fc29ca5b9a89c6f599b68@23.22.61.251:38182",
	"enode://f5336731fb657199dfdcadbe2a4a527fcfdb0816f38eee855512804e24d4d1a1bd9a3bfc1ec0fe9b5cd22001a018e7829ba6fc44d480356acfa331888df3a75b@35.173.125.30:38182",
	"enode://f7dc512940ca4a8f6858632abbdfc59cea6c4ed7a8da41ddfc4e4dac74e2664e74355fd7c688b285a22295e0053a800f759c9123ec741285a5bd602f89720cea@54.198.51.232:50674",
	"enode://0dc87959e0d4b9af89a9af7fe5dc9752f68ce40fd0d38537af34bb57d61cf4fd60cbefad3824c8725630c696560c6f8e88a5fb0801e4c4f3f7c2d0202fab8751@18.183.90.147:51300",
	"enode://3ab066e6807024a4514163e6688f7f5228a3bfec230b5e6c2a572b2d8081ad1d45a41d4172ab45a5dcf85f007c3c77afcd7d51496b96b0872567c1e4578e0760@54.194.101.180:57204",
	"enode://5df2f71ae6b2e3bb92f92badbce1f601feabd2d6ce899cf8265c39c38ff446136d74f5bfa089532c7074bb7606a509a54a2ac66397aaaab2363dad3f43c687a8@79.125.103.83:60752",
	"enode://a1d15b272cb4a427fc1c183d901926ea84151ec10c3d72fa2f1523b3f17b1ebc5093f2e3e37e8a3a98f3865ada3360892c6e880816cb8684161cb1ede5525183@142.132.192.59:37556",
	"enode://406d9615acca0d0b8fb3d8766ad51d73e4eee44f16d1a0d22388c028113bbeb662932689ebc559bb4f20ddba52b22f099baa176ee4d60fa859514eeb8e10d24c@23.88.72.56:43026",
	"enode://325e4a4d38099a56f89274a33da261617fa66517ebc9477a9f77691ad6f7d570f683d185b0af8d906060f84353f4ed5d9e5c7df9368a066f0446d30fdd253325@34.237.83.222:51060",
	"enode://59af33d546bf819ba42f66846accf48194336b9af72914e3523d57ab0c383f7ef11dc1f60f559eed8340240d1b19859b334a2bcbd0ebf0da1a81eb4cc39ec62c@3.237.191.210:37050",
	"enode://b0e5850bd5bf6421e39fdbffc5caf9fcec3f0a16464a43c75f255bfbdae52103520158fb137a781076b4f5f65dad4d51947c1ccd5cd1695e13ba6481dc08e9e2@3.112.109.205:59810",
	"enode://1a518939d25abd8c0e2397139385458b7987b45adfdfaa0ca18acfa1b7408d472305912d953f8fcf409e4adc5099c047fbcd61aa540deb24257b3ae73da51e1c@207.161.55.132:49586",
	"enode://3620495844fcc173163e53730d2172d72a63887f0779039396460fb40cabba83b5a6ed0e995932a38e8025c40b2edbec9bd340c55b7eaef6d745d472c2d3df79@54.197.158.138:42086",
	"enode://2305a46d20bc1d0f8be7e617b23b28f2fe460acc72e99d60128671e5eeffd4e4207d4646586d4b4953e88405cd324afaa0de9f5497e1be0fb7cc1ed3c0794960@44.199.89.100:51650",
	"enode://6b451954c885814136cc799db8f2f3f3d1d375ee3bde7bda2f93737a4900a9ff5158e6cfa419898eff008c0820be269ec28872e28a19f930cdf1fa74a1b10044@162.55.240.202:30311",
	"enode://a3717c4d788f1dc8c35126b02140a305417566df0d7ae7b8e7a724abd17ee872c473e3e943dcfc293a6fbc567fd8d0e41398938bf1475c6062652c0777202164@18.207.130.22:48092",
	"enode://8e66c3688ea92f41d94b1197a16f5d09b93d3c2a5bc1a8a3598b8ff0de755f2d9a9ea53f4fcff0281d0c416e7fb1ef52052926fc0fb75c9fc02beddd47e44fb4@54.152.241.156:37048",
	"enode://78c720b43ebe1e039eb14d70adb9a8b7ece30bf2c66962c60bfa566322c4598e4ef08decc449b98a02081a99054e20ac2cbd691b08cf8a17ade702431c2f152f@15.164.219.242:34494",
	"enode://f66f3194cc5996d4fcd823a24e335b455f724c4a39721b73e6e7bd5056f07f850764e231dc33af7abf843035792a2fdcd0132c59b3085837488e3fc1f92e6746@34.204.181.42:40942",
	"enode://16c7e98f78017dafeaa4129647d1ec66b32ee9be5ec753708820b7363091ceb310f575e7abd9603005e0e34d7b3316c1a4b6c8c42d7f074ed2eb4d073f800a03@3.85.216.212:38742",
	"enode://ccc50bf1db8e1eaa7013aad3a7fc76c004cbbdcb7fca2a9727ed4af2b71a92dc1a4b88621878158f38c7a412db561fb8a04e692554c40d69cb4964b1932b17c8@3.92.0.205:48936",
	"enode://124e4e34dccb65e699236e30e4749626fff6f10029e9dd82a2b56b3d317fb51c1f258be762d27ef1149adda45f074e1bc32c1f85b19347dc0f7018274520f5fe@13.115.249.240:52872",
	"enode://a6c4bf8cf50e7af74a52845d596bd77252a62f0f9377873306e22fbebe04a6e1395d006c5251a7dec7d2b32766f812b50ee3ca3a1c4b5f51f53870bec1d4a09c@3.250.227.125:51082",
	"enode://9e424a47df101e48ee5659e14b98d1130988a27ec37809e48fbf87768446cebd043976dcf296c69ed2afe9e4e96f62906d17386a4a46d070fc39a17a8a5e9e60@34.246.200.12:52084",
	"enode://f1e436fcb73d9b0ed48b56fd65a79c7fd58f1483b9424617d05e843c3cf1f5c6c34cf2114bafebdf39fdc96d8bddbaad0eec2f55f632512c0490654461c4b10d@34.245.28.18:49536",
	"enode://bb7e70a3d78e72ba7d4ae38d5a822bac1a99a8516841cbbbff3b32fd51dcf3e7b7e9fd04b705977f0de06aed40da9eddfb52ef01b20dc7a76c4831b39b818d17@65.21.123.68:60194",
	"enode://e514ba145a897c993d8b569d7bdd65f39f72ed03cb3b40882e70a95c2dfd30c3cf0e5707a4fff34ab7648132377dd012c54ac6035571a107b1548575ed1580b0@169.197.141.151:43080",
	"enode://8cbc21a3e28532a7de60a078cd5ad7dd292010a387c1372cb355fbb2332ae2cabb16b2b06dbbb932d7e7912d12f076a90d88c9a625729ad8bd951bbda336b985@34.244.8.134:43520",
	"enode://27fd54924ddd14688f348e83d9c34bdb8662b6e79595e38c9f97d9ce951ec9206d9b7f79aaf8b05127858dd25426a9e6b7421b056b2d25b29982f77cf342adfc@104.248.198.81:44712",
	"enode://47ae994b2c2190a1ed07f160573c038f364cd908cefad44357c8cab492d6d32b06662bd0d1844206a1907388c2adbd534174079da5f93046d9c8e2cedd85e494@142.132.157.156:43096",
	"enode://23c6aed5f69e81f49481fb00901f04f2fa21764f4cb8d13e395f9d538bd63626e6343387e832d3dc5b20a43a2104bb76221858845ca3e698bcba7e32a88dde6f@37.19.218.193:34736",
	"enode://9adfb0d75b622a9811a11f5e2e08bf65a9fa739b5e3ee516ce2f540eacb550ec89ab8fe6c42e24148a0b56bf940b4630049efd3f1defd2a4d0af0837bf13db62@34.204.30.113:41594",
	"enode://328449b52fc93508267e84c6b68c2d75a9543090888696d4ab7cc2201f796faa112a15897c89b03f199ff76677dd848dd427b12574b9eb25c6f8d8d5a7b63e63@54.83.228.97:49144",
	"enode://238d87c614b00629654df1d29e77b989011d30ebc822dee25d42231878c2d68606ee304c5828b19764bc0576bc7b0920ef86cb66e3adb5e002c02035906ca26f@3.238.239.247:47168",
	"enode://2814c79a0629a3d94280850a8a89cf94be788ea73c4d34b3c318211016571f55660f37933aa22b369d24b37c1f4a9828ada86f946bd6ece1900dde1f3eddcd57@47.251.4.101:30311",
	"enode://1881df2164b3f2326e0c9a9a15b2bf4c355d8ca625aaa6c38c32c8106fbe2d34de0d9fa77bda6c20e6eb2f43c94c576ede5394ce62580601d133ccc234b390bc@3.83.117.135:53562",
	"enode://17c4492a9d0f3ba04dad24e8d9c88b0fb20c94d412be64448045750aeaf69a2bcd4effdeb6ec89c65e019d4f8c213a0cf63ee3aca397378cc8cda726b3e43cf6@3.235.68.17:46070",
	"enode://9134d875f2be79e5ff08439da0c64ccbd6ac71539dec8568430c9e011c1a232addfa57fe2bdadaec1f52b681b6d5573d72a71bce3c094a3e3359bb91ea27af98@54.227.37.201:56004",
	"enode://b16e97ca25a8c9d4d43257519362cb3875bfc569e3c9a105b0fc0b813e26025c6a1e19a26d210c7130f9e48c81b6c87311bf496fbb3df56338766ff0d21320ca@107.20.3.143:42776",
	"enode://b25d3eeb43b1dafdbeeb0a9933b4db6750c6fe7c6aceed9a7b3c38fac4da729d06f7c742c2b6d708287eb195228d9007a7c9a7d93ec5ead3b2cd7163f3f97f3b@46.21.149.242:55476",
	"enode://bae08120adea77dba6533cfd43d3ae8f61ed49583d1a3ca2dabc9d142520a0432dea0ccebc08e6db2dbe848a8369dc498cf61050c60922140160fe734c302b39@18.232.230.87:44212",
	"enode://291d950ff1dbea7bb04b9cedce5d612c53e4b7002ac590ea86defa181f34680ee6e112d12a15e4cb46d3e65aa91d3863e4367a4435473a46e8732fc1adcaa8e1@54.74.49.155:50818",
	"enode://35b97bb92f07c3f6dfd1da515396cb613ec8c2ca1ef9d85826b11052ad217e48dc8565064332f72fb9a8e8c1218d386a2b948f8a60785a75f32bb79fa09e53c2@3.237.255.200:46036",
	"enode://a76805a9a6dfde2ad023ed8c33d4bff808b965f77b5e5d5c1971b1eac0c792ceaedae92bb28044bdd2edf308dc571b84d30e75888cf5606e29da6a3dc8850a0b@131.153.216.90:39838",
	"enode://9668c022915b307b5d7c30e11dd20f3aa19ddf180be7c20a45f641f417512cbb5452e9c39e60c50fc76093ee31cffac9fb4b6d205857f4cf4012f41c44f917bf@45.77.223.147:55146",
	"enode://4e7de268f922fec34745ec714f5848f1b4db1ac65132d0b24c714974272f977617228cb2ec64b8f12b701c5f496b41d8abfa1fd71de33a103b8147cfd730b780@54.159.161.144:50008",
	"enode://00d3375e8186b0c53432dcbf521b2c6e5eed9645e1d1a637c897cd0039ab4b3b697b93cd974530ec48be5c9d500e97f1daf103d738a98366146fecb84a607aee@44.200.101.218:50122",
	"enode://63292b5fd66c0640b14128ee359313ab8797861c785a91de80c9b5e7157cfbe91247ec14750b95878693c75c6381dd5587383c13fb1a5f4dedca617e3623ec8a@23.20.128.77:45358",
}

var BscStaticPeers = BscBootnodes

var ChapelBootnodes = []string{}
var ChapelStaticPeers = []string{
	"enode://69a90b35164ef862185d9f4d2c5eff79b92acd1360574c0edf36044055dc766d87285a820233ae5700e11c9ba06ce1cf23c1c68a4556121109776ce2a3990bba@52.199.214.252:30311",
	"enode://330d768f6de90e7825f0ea6fe59611ce9d50712e73547306846a9304663f9912bf1611037f7f90f21606242ded7fb476c7285cb7cd792836b8c0c5ef0365855c@18.181.52.189:30311",
	"enode://df1e8eb59e42cad3c4551b2a53e31a7e55a2fdde1287babd1e94b0836550b489ba16c40932e4dacb16cba346bd442c432265a299c4aca63ee7bb0f832b9f45eb@52.51.80.128:30311",
	"enode://0bd566a7fd136ecd19414a601bfdc530d5de161e3014033951dd603e72b1a8959eb5b70b06c87a5a75cbf45e4055c387d2a842bd6b1bd8b5041b3a61bab615cf@34.242.33.165:30311",
	"enode://604ed87d813c2b884ff1dc3095afeab18331f3cc361e8fb604159e844905dfa6e4c627306231d48f46a2edceffcd069264a89d99cdbf861a04e8d3d8d7282e8a@3.209.122.123:30311",
	"enode://4d358eca87c44230a49ceaca123c89c7e77232aeae745c3a4917e607b909c4a14034b3a742960a378c3f250e0e67391276e80c7beee7770071e13f33a5b0606a@52.72.123.113:30311",
}

var RialtoBootnodes = []string{}
var RialtoStaticPeers = ChapelStaticPeers

// ErigonBootnodes are the enode URLs of the P2P bootstrap nodes running on the ErigonNodes devnet
var ErigonBootnodes = []string{}

var SokolBootnodes = []string{
	"enode://f11a0f80939b49a28bf99581da9b351a592ec1504b9d32a7dfda79b36510a891e96631239c4166e5c73368c21e9bb3241e7fd6929b899772e5a8fe9a7b7c3af6@45.77.52.149:30303",
	"enode://e08adce358fc26dfbe1f24ee578dceaa29575ca44a39d9041203131db5135aceba6241840a9b57b1540eeaf7b4eff1aead28a74641be43342c35af454abb31b3@199.247.18.10:30313",
	"enode://f1a5100a81cb73163ae450c584d06b1f644aa4fad4486c6aeb4c384b343c54bb66c744aa5f133af66ea1b25f0f4a454f04878f3e96ee4cd2390c047396d6357b@209.97.158.4:30303",
	"enode://f11a0f80939b49a28bf99581da9b351a592ec1504b9d32a7dfda79b36510a891e96631239c4166e5c73368c21e9bb3241e7fd6929b899772e5a8fe9a7b7c3af6@45.77.52.149:30303",
}

var KovanBootnodes = []string{
	"enode://30499bde23362f7d310a34518a2a6ff765921870bf0c3e63d21153cfa7ba9cf39cc7c8e54e9dad2f2b3c07288b3e91b220656833cc2d843a54875c229f3f959a@8.9.8.175:30303",
	"enode://16898006ba2cd4fa8bf9a3dfe32684c178fa861df144bfc21fe800dc4838a03e342056951fa9fd533dcb0be1219e306106442ff2cf1f7e9f8faa5f2fc1a3aa45@116.203.116.241:30303",
	"enode://49a0e1aa38caa12cbf31222cb4e31cef1e8794cb4dd38012f84498ac867b19584e29bbf6d53201d7dfd3b5eb0998a4d908d096ed4ddb5f9102c623852cd331ec@54.87.247.5:30303",
}

var FermionBootnodes = []string{}

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

var MumbaiBootnodes = []string{
	"enode://320553cda00dfc003f499a3ce9598029f364fbb3ed1222fdc20a94d97dcc4d8ba0cd0bfa996579dcc6d17a534741fb0a5da303a90579431259150de66b597251@54.147.31.250:30303",
	"enode://f0f48a8781629f95ff02606081e6e43e4aebd503f3d07fc931fad7dd5ca1ba52bd849a6f6c3be0e375cf13c9ae04d859c4a9ae3546dc8ed4f10aa5dbb47d4998@34.226.134.117:30303",
}

var BorMainnetBootnodes = []string{
	"enode://0cb82b395094ee4a2915e9714894627de9ed8498fb881cec6db7c65e8b9a5bd7f2f25cc84e71e89d0947e51c76e85d0847de848c7782b13c0255247a6758178c@44.232.55.71:30303",
	"enode://88116f4295f5a31538ae409e4d44ad40d22e44ee9342869e7d68bdec55b0f83c1530355ce8b41fbec0928a7d75a5745d528450d30aec92066ab6ba1ee351d710@159.203.9.164:30303",
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
	case SepoliaGenesisHash:
		net = "sepolia"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case MumbaiGenesisHash:
		net = "mumbai"
	case BorMainnetGenesisHash:
		net = "bor-mainnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
