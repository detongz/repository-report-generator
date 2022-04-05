
# 1. 统计数据

共20位作者提交了36个 Commit 。感谢以下作者的贡献：

ruanwenjun,chenhu,CalvinKirs,zhongjiajie,BenJFan,kone-net,mans2singh,wuchunfu,yx91490,realdengziqi,kid-xiong,asdf2014,bigdataf,tobezhou33,bestcx,JesseAtSZ,asdf2014,CalvinKirs,chaozwn,yuangjiang

最近1周，共修改新增代码1096
行 ，删除代码3896行 。


# 2. 主要进展
## 2.1 新增功能

	- [Feature][Connector] update flink elasticsearch sink version
	https://github.com/apache/incubator-seatunnel/pull/1631

	- [Feature][Connector] Change ClickhouseFile read temp data method to local file
	https://github.com/apache/incubator-seatunnel/pull/1567

	- [Feature] [transform] To quickly locate the wrong SQL statement in flink sql transform
	https://github.com/apache/incubator-seatunnel/pull/1603

	- [Feature][core] Optimize the plugin load, rename plugin package name
	https://github.com/apache/incubator-seatunnel/pull/1592

	- [Feature] Support JDK 11 compiling
	https://github.com/apache/incubator-seatunnel/pull/1559

	- [Feature][Core]implament StructuredStreamingExecution
	https://github.com/apache/incubator-seatunnel/pull/1400

## 2.2 Bug 修复

	- [Bug] [core] Fix the resource leak for SparkStarter
	https://github.com/apache/incubator-seatunnel/pull/1655

	- [Bug][seatunnel-connector-flink-KafkaSink] fix not being able to send data to Kafka.
	https://github.com/apache/incubator-seatunnel/pull/1598

	- [Bug] [seatunnel-connector-spark-file] fix file source and sink default params
	https://github.com/apache/incubator-seatunnel/pull/1590

	- [Bug] [core] Arbitrary file write during archive extraction ("Zip Slip")
	https://github.com/apache/incubator-seatunnel/pull/1583

## 2.3 功能改进

	- [Improve][Dependency]Unified version management
	https://github.com/apache/incubator-seatunnel/pull/1642

	- [Improve][core] add Override on getDeployMode(#1637)
	https://github.com/apache/incubator-seatunnel/pull/1638

	- [Improvement][flink-connector-file] Updated default rollover interval constant name
	https://github.com/apache/incubator-seatunnel/pull/1616

	- [Improve][Doc] Add kafka source parameter explanation
	https://github.com/apache/incubator-seatunnel/pull/1624

	- [Improvement][build] optimize package name
	https://github.com/apache/incubator-seatunnel/pull/1610

	- [Improvement][core] Rename seatunnel-core-sql to seatunnel-core-flink-sql
	https://github.com/apache/incubator-seatunnel/pull/1609

	- [Improvement][flink-connector-kafka] rename kafkasink/add semantic/simplify code
	https://github.com/apache/incubator-seatunnel/pull/1607

	- [Improve]Some code block optimizations
	https://github.com/apache/incubator-seatunnel/pull/1600

	- [Improve][Connectors]Refactor es config class
	https://github.com/apache/incubator-seatunnel/pull/1594

	- [Improve][Connector]Exceptions should be either logged or rethrown but not both
	https://github.com/apache/incubator-seatunnel/pull/1587

	- [Improve] [core] Must throw exception while creating plugin instances
	https://github.com/apache/incubator-seatunnel/pull/1419

## 2.4 WIP

	- [Draft] demo flink mysql cdc
	https://github.com/apache/incubator-seatunnel/pull/1362

	- [Bug] [test] Auto close FileSource
	https://github.com/apache/incubator-seatunnel/pull/1634

	- [Improve] [api] Improve sink api by removing useless extends Object
	https://github.com/apache/incubator-seatunnel/pull/1635

	- [Improve] [core] Remove useless dependencies
	https://github.com/apache/incubator-seatunnel/pull/1602

	- [CI]Merged PR should be after basic CI check passes
	https://github.com/apache/incubator-seatunnel/pull/1620

	- [Refactor] [core] Remove the uselss return value from outputBatch of FlinkBatchSink
	https://github.com/apache/incubator-seatunnel/pull/1462

	- [Improve] [core] Enhance the cd logic
	https://github.com/apache/incubator-seatunnel/pull/1571

	- [Improve] [ci] Upgrade calcite-druid from 1.29.0 to 1.30.0
	https://github.com/apache/incubator-seatunnel/pull/1612

	- [Feature] [connnector] Apache HBase Flink Connector
	https://github.com/apache/incubator-seatunnel/pull/995

	- [Feature][Metric]seatunnel spark job metric
	https://github.com/apache/incubator-seatunnel/pull/1578

	- [Improve] [core] Simplify the check logic of directory
	https://github.com/apache/incubator-seatunnel/pull/1585

	- [Feature] [transform]  Add concat transform for Spark plugins
	https://github.com/apache/incubator-seatunnel/pull/1166

## 2.5 其他

	- [hotfix][directory] recover plugins directory
	https://github.com/apache/incubator-seatunnel/pull/1651

	- [DOCS]Update spark dependency version notes
	https://github.com/apache/incubator-seatunnel/pull/1627

	- [Dependency]Upgrade log4j-core to 2.17.1
	https://github.com/apache/incubator-seatunnel/pull/1646

	- [License]Redeclare dual license to avoid ambiguity
	https://github.com/apache/incubator-seatunnel/pull/1644

	- [ci] Run auto license only if commit message contains keyword
	https://github.com/apache/incubator-seatunnel/pull/1643

	- [hotfix][connector-redis][docs] Updated required value in attribute tables for those with default values
	https://github.com/apache/incubator-seatunnel/pull/1629

	- [Chore]Modify irregular code
	https://github.com/apache/incubator-seatunnel/pull/1633

	- [doc] Fix dead link to start
	https://github.com/apache/incubator-seatunnel/pull/1618

	- [donot-merge][ci] Migrate check-license into license.yml
	https://github.com/apache/incubator-seatunnel/pull/1621

	- [Hotfix][ci] Sort dependencies to pass the ci in known-dependencies.txt
	https://github.com/apache/incubator-seatunnel/pull/1453

	- [hotfix][seatunnel-spark-examples][#1614] lose default deploy mode value issue
	https://github.com/apache/incubator-seatunnel/pull/1615

	- [doc] Add hint only spark supported for Mesos cluster
	https://github.com/apache/incubator-seatunnel/pull/1617

	- [Imporvement][directory] Remove plugins directory
	https://github.com/apache/incubator-seatunnel/pull/1613

	- [doc] Change docs structure and make it more suitable
	https://github.com/apache/incubator-seatunnel/pull/1611

	- [SeaTunnel#1218] rewrite spark start script with java
	https://github.com/apache/incubator-seatunnel/pull/1486

	- [hotfix][hbase][docs] Corrected staging_dir attribute
	https://github.com/apache/incubator-seatunnel/pull/1588

	- [doc] Move sidebars.js from website repo
	https://github.com/apache/incubator-seatunnel/pull/1586

