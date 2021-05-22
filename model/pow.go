package model

import (
	"blockDemo/util"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//区块链的重要设想就是如果你要往里面添加新的区块就要完成一些艰难的工作。
//而正是这种机制确保了区块链的安全和数据的一致性。
//同时，给这些艰难的工作适当的奖励（这也是人们挖矿获取比特币的机制）。
//
//这种机制与现实非常类似：一个人必须通过努力工作获得回报以维持生计。
//在区块链当中，网络上的参与者（矿工）的工作维持网络的正常运行，向区块链中加入新的区块，并因为他们的努力工作而获得回报。
//他们的工作结果是将一个个区块以安全的方式连成一个完整的区块链，这也维护了整个区块链数据库的稳定性。
// 更有价值的是，谁完成了工作必须进行自我证明。
//
//这一整个的“努力工作并证明”的机制被称为“工作证明”（PoW）。
//它难在需要大量的计算资源：即便是高性能的计算机，也无法快速完成工作。
//甚至，为了保证新的区块增加速度维持在6个每小时，这个计算工作会越来越繁重。
//在比特币当中，计算工作的目的是为了给区块找一个匹配的并满足一些特定要求的哈希值。
//同时这个哈希值也为工作服务。因此，实际的工作就是寻找证明。
//
//最后一点需要注意的，PoW算法必须满足一项要求：
//虽然计算很困难，但是工作证明的验证要非常容易。
//因为证明通常会传给网络上的其他参与者进行，不应该消耗他们的太多时间了验证这个证明。
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

const targetBits  = 24
var (
	maxNonce = math.MaxInt64
)
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			util.IntToHex(pow.block.Timestamp),
			util.IntToHex(int64(targetBits)),
			util.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}