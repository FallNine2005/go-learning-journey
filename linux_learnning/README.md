|第一天的学习内容|

"使用linux系统成功的创建了目录和文件，学会了vim，touch，rm，mkdir等指令的操作。"
"成功使用cp指令更新了目录结构，修复了git库。"
"创建了mv1.go文件并使用mv指令改名为mv.go。"
"用cat指令查看README.md。"
"建立了big_file.txt并用less功能交互。"
"学会了head，tail，echo的使用。"
"使用tail -f实现了对bigfile的追踪，并用echo和>\>>实现了输出转移并追踪。"
"使用ls -l >实现了将文件目录覆盖bigfile文件，并学习了cat >，echo >>等指令。"
"cal指令显示日期并覆盖bigfile"
"运用ln -s指令建立了软连接，用test_txt指向了bigfile。"
"history"

|第二天的学习内容——接下来同步github会在一天的内容完结时一同发送|

“学会了使用date指令查看时间，cal指令查看年历。”
“学会使用find指令查找并将结果写入bigfile。”
“使用updatedb完成了遍历，并用locate查找”
“使用| gerp完成内容“March”的查找和显示行号查找Mo并写入bigfile”
“使用gzip对bigfile压缩并用gunzip解压”
“用zip -r将整个linux learning文件夹压缩为testzip.zip”
“使用tar -czvf将01，02，linuxlearning文件夹打包压缩为resttar.tar.zip，随后解压到新建testtar文件夹内”
“使用ls -ahl 查看文件的组和所有者，使用useradd -g A testA来建立了一个属于组A的测试账号testA，切换账号并生成了一个test.txt文件，查询到其属于组A，使用chgrp将其组别改为B。使用usermod将testa改为组B再使用chgrp指令将testa家目录下的所有文件改为组B。”

|第三天的学习内容|

“使用ll指令查询文件的权限并打印去bigfiles，学习了，-，l，d，c，b的对应表示内容，学会了0-9位前置字符代表的权限含义，使用chmod修改bigfile权限为rwxr-xr--。”
“学会了chown改变所有者和组的用法，将testa的文件夹所有者改为fallnine2005”
“学会了chgrp（先前已经使用过）”
“完成了警察土匪游戏实例操作————设置C，D组，设置用户police，bandit并分别加入C,D组，使用police创立文件bandit_will_die，设置权限640，随后修改为764，再将土匪bandit投降从D组移入C组检验是否可以查看警察police的文件。结果：成功。并写入bandit success to read,write this file。”
“创建了testcrontab.txt文件，并使用crontab指令每天调用一次文件。每天会在下午2点时自动写入read.me。
创建了testcrontab.sh脚本，并每天下午2点自动写入当前的时间日期进testcrontab2.txt文件中。”
“使用ps -ef | gerp atd来检测atd进程是否在运行，写入一个测试用at命令，将会在5点写入一个一次性测试语句。”Sun Nov  1 17:00:00 CST 2026
succeed to use at

|第四天的内容|

"使用lsblk指令查看了硬盘的挂载情况。"
“学习了fdisk添加硬盘和分区的相关知识————格式化mkfs -t ext4,挂载mount等，但在使用wsl系统时无法复现教程内容，暂且挂起。”
“使用df -h查询分区空间”
“使用du -h查询磁盘空间,使用du -hac -d1查询目录的使用情况。”
“用ls -l | grep "^-" | wc -l查询目录下普通文件的数量。”
"用tree来以树状图显示目录，没有指令可以用yum install tree下载安装"
"使用ifconfig查询wsl的虚拟ip，在主机终端管理员处用ipconfig查询到主机ip，学习了vm虚拟机和linux内网及主机外网的关系，wsl系统与vm虚拟机网络结构并不相同。"
"在主机设置了镜像模式，使得wsl能直接通过主机联通外网，wsl现在可以直接解析域名然后上网了。"
"理解了如何在文件中手动修改vm虚拟机中linux系统的ip，网关"
“使用hostname查看主机名并尝试修改为fallnine2005-pc,但wsl会强制同步主机的pc名，因此失败，暂时挂起。”
"镜像模式出现了问题，导致22端口被占用无法连接MoBaX，取消了镜像模式。"
“学习了主机名与ip的关系，尝试了主机名设置与映射，但是因为wsl的特性暂时无法实验，暂且挂起。”

|第五天的学习内容|

“开始学习程序与进程（前台后台），学习了ps -aux指令的运用”
“使用ps -ef显示父进程，使用kill通过进程号中止进程，使用killall通过进程名中止进程，选项-9用来强行结束。”
“使用pstree显示进程的进程树，使用setup查看所有系统服务进程，使用service管理后台进程守护进程等，但是wsl环境不支持setup，转而使用service --status-all。”
"可以通过使用systemctl get-default来查看当前运行级别，使用systemctl set-default TARGET.target来设置默认启动运行级别."
“使用chkconfig命令来设置服务的运行级别的自启动，wsl中暂时没有可以被chkconfig管理的服务。”
“学习了systemctl命令来管理服务的相关知识。包括查看服务状态（status）设置服务的自启动（enable/disable），查看服务的自启动状态（list-unit-files）等。”
“wsl环境并不包含完整的systemctl功能，因此相关实践暂且挂起。”
“学习了firewall-cmd --permanent --add-port/--remove-port防火墙指令来控制端口的开关。用netstat -anp来查看端口的协议。”
“学习了使用top指令来监控正在进行的进程（动态监控），使用u来监控特定的用户，使用k来结束进程。”
“学习了用netstat监控网路状态，知道了tcp的超时等待功能。”

|第六天的学习内容|

“学习了RPM相关的知识，使用rpm -qa来查询已经安装的rpm包，rpm -q来查询某个软件是否安装，qi查询软件包信息，ql查询软件包内容，qf查询路径文件属于哪个文件包，-e用来删除rpm包”
“使用rpm -ivh来安装rpm包，i安装v提示h进度条。”
“学习了用yum list来查询已经安装的rpm包，使用yum install来下载软件，-e删除软件”
“wsl环境暂不推荐使用rpm和yum来下载管理软件，相关实践暂且挂起。”
“下载了mysql数据库以进行配套操作。”
“学习了和脚本shell，bash的相关知识，脚本需要以#！/bin/bash开头，脚本需要有执行权限，建立了testshell文件夹。”
“写入文件hello.sh,赋予754权限，执行输出hello world！使用相对路径或绝对路径执行，或用sh+脚本执行。”
“使用echo查询系统变量如$PATH.”
"使用set查询所有的系统变量"
“使用变量=值来设置变量，unset来删除变量，shell中不能存在空格，readonly可以设置静态变量。”
“新建var.sh用于测试变量”
“使用A=``反引号命令将返回值赋值给A，等价于A=$()”
“使用export将shell变量设置为环境/全局变量。使用source来使变量生效。因此可以在管理环境变量的文件中加入export 变量名=变量值来将一些变量变为全局变量。然后通过echo $变量名来查询变量值。”
“使用：<<！内容！来多行注释。”
“学习了位置参数变量的相关知识。”
“学习了预定参数的相关知识。$$用来查询当前进程的pid，$!查询后台的最后一个进程的pid，$?查询上一个进程是否执行完毕。”
“可以用&来使进程在shell后台运行。”
“学习了运算符相关的知识，可以在$(())或$[]中使用。EXPR写法中\*乘 /除 %取余。”

|第七天的学习内容|

“学习了shell编程的判断语言[ condition ],是真返回True，非真返回False。[ condition ] &&  echo ok || echo notok,条件满足执行后续否则执行额外。”
“判断常用语法。-lt小于 -gt大于 -le小于等于 -ge大于等于 -eq等于 -ne不等于。-r-w-x分别判断是否有读写执行的权限。-f -e -d分别判断是否是常规文件是否存在是否是文件夹。”
“编写了预输入变量脚本condition.sh，判断$1是否等于100并输出yes，成功。”
“学习了多分支相关的知识，新建if.sh文件。并设立了多分支判断的测试用编程内容。”
“学习了流程控制相关的内容，新建case.sh用于实践。”
“学习了for循环相关的内容，新建for.sh用于实践。使用100 200 300的循环数将num依次赋值给if.sh并充当$1执行输出。使用i=1；i<10；i++的条件迭代相加输出sum。”
“学习了while循环相关的内容，新建while.sh用于实践。”
"学习了read读取控制台输入。新建read.sh。"
“学习了shell编程的函数相关内容。使用函数basename获取了big_file字符串。还有与basename功能相反的dirname”
“学习了自定义函数[ function ]功能，新建function.sh。”
“学习了使用函数管理数据库，在wsl环境中下载了mysql数据库，用root用户登陆并创建了一般用户fallnine2005，并设置密码给予权限。”
“牵扯到数据库相关的知识，决定了在学习go语言的时候再学习数据库，因此相关的实践暂且挂起。”

|第八天学习的内容|

“开始学习和日志相关的知识。”
“加入了中文包，修改wsl为中文环境。”
“因为wsl环境不会默认开启rsyslogd日志管理服务，这里设置为手动开启。”
“wsl环境的目录内容与正常linux系统的目录并不相同，没有传统意义上的boot和secure等日志文件。”
“学会了查看日志。鉴于linux环境下的go语言使用情况，暂不学习如何管理日志服务，等go语言使用时需要再学习。”
“学习了系统权限划分相关的知识。”
"学习了系统计量和访问相关知识。"
“关于linux的学习暂时告一段落，接下来将会在golang语言的学习过程中继续磨合学习。”