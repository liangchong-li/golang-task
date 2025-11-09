package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/victor/ethclient/task/counter"
)

func main() {
	// task1()
	// task2()
	onlyRead()
}

// 查询区块
// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
// 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
// 输出查询结果到控制台。
// 发送交易
// 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
// 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
// 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
// 对交易进行签名，并将签名后的交易发送到网络。
// 输出交易的哈希值。
func task1() {
	client, err :=
		ethclient.Dial("https://sepolia.infura.io/v3/e98785ce63ef4bf9a8b977697a83e786")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("时间戳: %d\n", block.Time())
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))

	// 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
	// 对交易进行签名，并将签名后的交易发送到网络。
	// 输出交易的哈希值。
	privateKey, err :=
		crypto.HexToECDSA("e3850c6fa1e1d7fe9adbb98b126f1a97b3508147696fd4850f2fc1d6d346548f")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// fmt.Printf("公钥: %v\n", fromAddress)
	// fmt.Printf("公钥: %v\n", "0x86a7cf6f44effcd89041786dafceda1672f7a253")
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // 0.01 ether
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x7c7c11d08b95c76fa5aa63d51eb3c0f8646cd57f") // 替换为接收方地址

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易哈希: %s\n", signedTx.Hash().Hex())
}

// 使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
//
//	具体任务
//
// 编写智能合约
// 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
// 编译智能合约，生成 ABI 和字节码文件。
// 使用 abigen 生成 Go 绑定代码
// 安装 abigen 工具。
// 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
// 使用生成的 Go 绑定代码与合约交互
// 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
// 调用合约的方法，例如增加计数器的值。
// 输出调用结果。
func task2() {
	// 1. 全局安装 solc
	// npm install -g solc
	// 2. 生成 ABI 和字节码文件
	// solcjs --bin Counter.sol
	// solcjs --abi Counter.sol
	// 3. 安装 abigen 工具
	// go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	// 4. 使用 abigen 生成绑定 go 代码
	// abigen --bin=Counter_sol_Counter.bin --abi=Counter_sol_Counter.abi --pkg=counter --out=counter.go

	// 5. 部署合约
	client, err :=
		ethclient.Dial("https://sepolia.infura.io/v3/e98785ce63ef4bf9a8b977697a83e786")
	if err != nil {
		log.Fatal(err)
	}
	private, err :=
		crypto.HexToECDSA("e3850c6fa1e1d7fe9adbb98b126f1a97b3508147696fd4850f2fc1d6d346548f")
	if err != nil {
		log.Fatal(err)
	}
	public := private.Public()
	publicKeyECDSA, ok := public.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err :=
		bind.NewKeyedTransactorWithChainID(private, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("合约地址: %s\n", address.Hex())
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())

	// 6. 使用合约
	opt, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{Context: context.Background()}

	tx, err = instance.Count(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("first do: ", tx.Hash().Hex())
	number, err :=
		instance.Number(callOpts)
	if err != nil {
		log.Fatal(err)
	}
	// 并不能立即得到最新的结果。交易被打包很慢
	fmt.Println("after first count number is: ", number)
	tx, err = instance.Count(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("second do: ", tx.Hash().Hex())
	number, err =
		instance.Number(callOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after second count number is: ", number)
}

func onlyRead() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/e98785ce63ef4bf9a8b977697a83e786")
	if err != nil {
		log.Fatal(err)
	}
	contractAddr := "0x636E29dC8C8E7280090DCc36e56e79BfC7B1fcFA"
	counterContract, err := counter.NewCounter(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("e3850c6fa1e1d7fe9adbb98b126f1a97b3508147696fd4850f2fc1d6d346548f")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	callOpts := &bind.CallOpts{Context: context.Background()}

	tx, err := counterContract.Count(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("first do: ", tx.Hash().Hex())
	number, err :=
		counterContract.Number(callOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after first count number is: ", number)
	tx, err = counterContract.Count(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("second do: ", tx.Hash().Hex())
	number, err =
		counterContract.Number(callOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after second count number is: ", number)
}
