# atb
convert alipay bill to beancount version

# 基本概念

参见 example.bean 中的实例，在 fava 中 equity 也就是净资产显示为负数。同时，Liabilities 也就是债务也显示为负数。
而 Income 总是一个递减的负数。

# 软件用法

下载代码编译，参见 ./atb -h

# 注意
支付宝账单中，手机上面添加的注释和 web 端并不通用。而导出账单的数据是根据 web 端来的。在 web 端添加注释不太方便，故不支持依据注释进行判断应该
进入哪个账号，而添加了依据交易对方名字和金额判断。
