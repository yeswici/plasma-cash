package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ChildChainService child client to reference server
type ChildChainService struct {
	url string
}

func (c *ChildChainService) CurrentBlock() (error, *Block) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/block", c.url), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("CurrentBlock Body:", string(body))
	block := Block{blockId: string(body)}
	return nil, &block
}

func (c *ChildChainService) BlockNumber() int {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocknumber", c.url), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("BlockNumber Body:", string(body))

	i, _ := strconv.Atoi(string(body))
	return i
}

func (c *ChildChainService) Block(blknum int) (error, *Block) {
	result := fmt.Sprintf("%s/block/%s", c.url, strconv.Itoa(blknum))
	fmt.Println(result)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/block/%s", c.url, strconv.Itoa(blknum)), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Block Body:", string(body))
	block := Block{blockId: string(body)}
	return nil, &block
}

func (c *ChildChainService) Proof(blknum int, uid int) (error, *Proof) {
	result := fmt.Sprintf("%s/proof/?blknum=%s&uid=%s", c.url, strconv.Itoa(blknum), strconv.Itoa(uid))
	fmt.Println(result)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/proof?blknum=%s&uid=%s", c.url, strconv.Itoa(blknum), strconv.Itoa(uid)), nil)
	if err != nil {
		fmt.Print(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Proof Body:", string(body))
	proof := Proof{proofstring: string(body)}
	return nil, &proof
}

func (c *ChildChainService) SubmitBlock(*Block) error {
	return nil
}

func (c *ChildChainService) SendTransaction() error {
	return nil
}

func NewChildChainService(url string) ChainServiceClient {
	return &ChildChainService{url}
}