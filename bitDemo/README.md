### <font color="red">比特币的总发行数量是怎么来的</font>

>- **比特币的发行规则**
>
>  - 每21万个区块奖励减半
>  - 开始奖励为50个比特币
>  - 知道奖励衰减为0，比特币发行完毕
>
>-  **根据比特币的发行规则，可以求出发行总数**
>
>  - 代码如下：
>
>    ``` go
>    total := 0.0 	//总的发行量
>    blocks := 21.0  //每次衰减时，产生区块的数量，单位：万个
>    reward := 50.0	//初始的挖矿奖励
>    for reward > 0 {
>    	sum := reward * blocks
>    	reward *= 0.5
>    	total += sum
>    }
>    ```

**执行结果，如下：**

![](https://github.com/AlexBruceLu/BlockChain/wiki/blockChain00.png)