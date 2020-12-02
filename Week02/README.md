# 学习比较
## 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

- 任何异常必须处理，且只处理一次
- sql.ErrNoRows在dao层抛出，到了service层，有三种做法
1、非透明的往上抛
2、errors.Wrap之后上抛
3、降级处理

在根据用户名查用户的场景下，service层发现数据查不到，可以认为这个实际场景中，是正常业务，并直接返回 error=nil,user=nil