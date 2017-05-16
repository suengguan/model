package model

import (
	// "config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	// 用户权限
	USER_AUTHORITY_ADMIN = 0x01 // 操作系统资源
	USER_AUTHORITY_USER  = 0x02 // 普通用户

	// log类型 1 ~ 100
	LOG_TYPE_DEBUG = 0x01 // 调试
	LOG_TYPE_INFO  = 0x02 // 正常运行过程
	LOG_TYPE_WARN  = 0x03 // 潜在错误，警告
	LOG_TYPE_ERROR = 0x04 // 错误事件，不影响系统运行
	LOG_TYPE_CRASH = 0x05 // 严重错误

	// job状态 101 ~ 200
	JOB_STATUS_TYPE_RUN    = 101
	JOB_STATUS_TYPE_STOP   = 102
	JOB_STATUS_TYPE_PAUSE  = 103
	JOB_STATUS_TYPE_FINISH = 104
	JOB_STATUS_TYPE_ERROR  = 105
	JOB_STATUS_TYPE_READY  = 0

	// table name
	TN_USER      = "USER_T"
	TN_PROJECT   = "PROJECT_T"
	TN_JOB       = "JOB_T"
	TN_MODULE    = "MODULE_T"
	TN_POD       = "POD_T"
	TN_LOG       = "LOG_T"
	TN_ALGORITHM = "ALGORITHM_T"
	TN_RESOURCE  = "RESOURCE_T"
	TN_ACTION    = "ACTION_T"
)

type User struct {
	Id                int64      `json:"id" orm:"column(ID)"`
	Name              string     `json:"name" orm:"column(NAME)"`
	Header            string     `json:"header" orm:"column(HEADER)"`
	Email             string     `json:"email" orm:"column(EMAIL)"`
	Phone             string     `json:"phone" orm:"column(PHONE)"`
	Company           string     `json:"company" orm:"column(COMPANY)"`
	EncryptedPassword string     `json:"encryptedPassword" orm:"column(ENCRYPTED_PASSWORD)"`
	CreatedAt         int64      `json:"createdAt" orm:"column(CREATED_AT)"`
	UpdatedAt         int64      `json:"updatedAt" orm:"column(UPDATED_AT)"`
	Active            bool       `json:"active" orm:"column(ACTIVE)"`
	Role              int        `json:"role" orm:"column(ROLE)"`
	Resource          *Resource  `json:"resource" orm:"reverse(one)"`
	Projects          []*Project `json:"projects" orm:"reverse(many)"`
	Logs              []*Log     `json:"logs" orm:"reverse(many)"`
	Actions           []*Action  `json:"actions" orm:"reverse(many)"`
}

func (this *User) TableName() string {
	return TN_USER
}

type Resource struct {
	Id                  int64   `json:"id" orm:"column(ID)"`
	AlgorithmResource   string  `json:"algorithmResource" orm:"column(ALGORITHM_RESOURCE);type(text)"`
	CpuTotalResource    float64 `json:"cpuTotalResource" orm:"column(CPU_TOTAL_RESOURCE)"`
	CpuUsageResource    float64 `json:"cpuUsageResource" orm:"column(CPU_USAGE_RESOURCE)"`
	CpuUnit             string  `json:"cpuUnit" orm:"column(CPU_UNIT)"`
	MemoryTotalResource float64 `json:"memoryTotalResource" orm:"column(MEMORY_TOTAL_RESOURCE)"`
	MemoryUsageResource float64 `json:"memoryUsageResource" orm:"column(MEMORY_USAGE_RESOURCE)"`
	MemoryUnit          string  `json:"memoryUnit" orm:"column(MEMORY_UNIT)"`
	User                *User   `json:"user" orm:"column(USER_ID);rel(one)"`
	QuotaNamespace      string  `json:"quotaNamespace" orm:"column(QUOTA_NAMESPACE)"`
	QuotaName           string  `json:"quotaName" orm:"column(QUOTA_NAME)"`
}

func (this *Resource) TableName() string {
	return TN_RESOURCE
}

type Project struct {
	Id          int64  `json:"id" orm:"column(ID)"`
	Name        string `json:"name" orm:"column(NAME)"`
	CreatedBy   int64  `json:"createdBy" orm:"column(CREATED_BY)"`
	CreatedAt   int64  `json:"createdAt" orm:"column(CREATED_AT)"`
	Description string `json:"description" orm:"column(DESCRIPTION)"`
	User        *User  `json:"user" orm:"column(USER_ID);rel(fk)"` //设置一对多关系
	Jobs        []*Job `json:"jobs" orm:"reverse(many)"`
}

func (this *Project) TableName() string {
	return TN_PROJECT
}

type Job struct {
	Id          int64     `json:"id" orm:"column(ID)"`
	Name        string    `json:"name" orm:"column(NAME)"`
	Description string    `json:"description" orm:"column(DESCRIPTION)"`
	CreatedBy   int64     `json:"createdBy" orm:"column(CREATED_BY)"`
	CreatedAt   int64     `json:"createdAt" orm:"column(CREATED_AT)"`
	Project     *Project  `json:"project" orm:"column(PROJECT_ID);rel(fk)"` //设置一对多关系
	Modules     []*Module `json:"modules" orm:"reverse(many)"`
}

func (this *Job) TableName() string {
	return TN_JOB
}

type Module struct {
	Id          int64  `json:"id" orm:"column(ID)"`
	Name        string `json:"name" orm:"column(NAME)"`
	InputFiles  string `json:"inputFiles" orm:"column(INPUT_FILES)"`
	OutputFiles string `json:"outputFiles" orm:"column(OUTPUT_FILES)"`
	Parameters  string `json:"parameters" orm:"column(PARAMETERS)"`
	Algorithm   string `json:"algorithm" orm:"column(ALGORITHM)"`
	Description string `json:"description" orm:"column(DESCRIPTION)"`
	Job         *Job   `json:"job" orm:"column(JOB_ID);rel(fk)"` //设置一对多关系
	Pods        []*Pod `json:"pods" orm:"reverse(many)"`
}

func (this *Module) TableName() string {
	return TN_MODULE
}

type Pod struct {
	Id        int64   `json:"id" orm:"column(ID)"`
	Name      string  `json:"name" orm:"column(NAME)"`
	PodName   string  `json:"podName" orm:"column(POD_NAME)"`
	RcName    string  `json:"rcName" orm:"column(RC_NAME)"`
	SvcName   string  `json:"svcName" orm:"column(SVC_NAME)"`
	DataRange string  `json:"dataRange" orm:"column(DATA_RANGE)"`
	Cpu       float64 `json:"cpu" orm:"column(CPU)"`
	Memory    float64 `json:"memory" orm:"column(MEMORY)"`
	Module    *Module `json:"module" orm:"column(MODULE_ID);rel(fk)"` //设置一对多关系
}

func (this *Pod) TableName() string {
	return TN_POD
}

type Algorithm struct {
	Id          int64  `json:"id" orm:"column(ID)"`
	Name        string `json:"name" orm:"column(NAME)"`
	Version     string `json:"version" orm:"column(VERSION)"`
	Description string `json:"description" orm:"column(DESCRIPTION)"`
	Author      string `json:"author" orm:"column(AUTHOR)"`
	Parameters  string `json:"parameters" orm:"column(PARAMETERS)"` // p1=1,p2=...pn=n
	Image       string `json:"image" orm:"column(IMAGE)"`
	Downloads   int64  `json:"downloads" orm:"column(DOWNLOADS)"`
	Stars       int64  `json:"stars" orm:"column(STARS)"`
}

func (this *Algorithm) TableName() string {
	return TN_ALGORITHM
}

// ========================================= log
/*
docker 原生支持的log 通用格式为
------------------
| time | content |
------------------
*/

type Log struct {
	Id        int64  `json:"id" orm:"column(ID)"`
	Time      int64  `json:"time" orm:"column(TIME)"`
	SessionId string `json:"sessionId" orm:"column(SESSION_ID)"`
	User      *User  `json:"user" orm:"column(USER_ID);rel(fk)"` //设置一对多关系
	DevType   int    `json:"devType" orm:"column(DEV_TYPE)"`
	Type      int    `json:"type" orm:"column(TYPE)"`
	Content   string `json:"content" orm:"column(CONTENT)"` // log内容
}

func (this *Log) TableName() string {
	return TN_LOG
}

type Action struct {
	Id        int64  `json:"id" orm:"column(ID)"`
	Time      int64  `json:"time" orm:"column(TIME)"`
	SessionId string `json:"sessionId" orm:"column(SESSION_ID)"`
	User      *User  `json:"user" orm:"column(USER_ID);rel(fk)"` //设置一对多关系
	DevType   int    `json:"devType" orm:"column(DEV_TYPE)"`
	Type      int    `json:"type" orm:"column(TYPE)"`
	Content   string `json:"content" orm:"column(CONTENT);type(text)"`
}

func (this *Action) TableName() string {
	return TN_ACTION
}

type PodLog struct {
	ProjectId   int64  `json:"projectId"`
	ProjectName string `json:"projectName"`
	JobId       int64  `json:"jobId"`
	JobName     string `json:"jobName"`
	ModuleId    int64  `json:"moduleId"`
	ModuleName  string `json:"moduleName"`
	PodId       int64  `json:"podId"`
	PodName     string `json:"podName"`
	LogContent  string `json:"logContent"`
}

type JobStatus struct {
	UserName       string          `json:"userName"`
	ProjectId      int64           `json:"projectId"`
	ProjectName    string          `json:"projectName"`
	JobId          int64           `json:"jobId"`
	JobName        string          `json:"jobName"`
	JobDescription string          `json:"jobDescription"`
	Status         int             `json:"status"` // run, stop, pause,finish
	Progress       int             `json:"progress"`
	ModulesStatus  []*ModuleStatus `json:"modulesStatus"`
}

type ModuleStatus struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      int          `json:"status"` // run, stop, pause,finish
	Progress    int          `json:"progress"`
	PodsStatus  []*PodStatus `json:"podsStatus"`
}

type PodStatus struct {
	Id       int64  `json:"id"`
	Status   int    `json:"status"` // run, stop, pause,finish
	Progress int    `json:"progress"`
	Reason   string `json:"reason"`
}

type Summary struct {
	ProjectTotalCount int     `json:"projectTotalCount"`
	JobTotalCount     int     `json:"jobTotalCount"`
	JobRunningCount   int     `json:"jobRunningCount"`
	JobSuccessCount   int     `json:"jobSuccessCount"`
	JobFailedCount    int     `json:"jobFailedCount"`
	CpuTotal          float64 `json:"cpuTotal"`
	CpuUsed           float64 `json:"cpuUsed"`
	CpuUsedPercent    int     `json:"cpuUsedPercent"`
	CpuUnit           string  `json:"cpuUnit"`
	MemoryTotal       float64 `json:"memoryTotal"`
	MemoryUsed        float64 `json:"memoryUsed"`
	MemoryUsedPercent int     `json:"memoryUsedPercent"`
	MemoryUnit        string  `json:"memoryUnit"`
}

type Email struct {
	Id      int64  `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ====================================================================================== kubernetes model
// ========================================= kubernetes model - const value
const (
	// version
	KUBE_API_VERSION = "v1"

	// resource type
	KUBE_RESOURCE_NAMESPACE             = "Namespace"
	KUBE_RESOURCE_REPLICATIONCONTROLLER = "ReplicationController"
	KUBE_RESOURCE_SERVICE               = "Service"

	// for replication controller
	KUBE_RC_RESTART_POLICY_ALWAYS    = "Always"
	KUBE_RC_RESTART_POLICY_NEVER     = "NEVER"
	KUBE_RC_RESTART_POLICY_ONFAILURE = "OnFailure"

	KUBE_RC_DNS_POLICY_DEFAULT      = "Default"
	KUBE_RC_DNS_POLICY_CLUSTERFIRST = "ClusterFirst"

	// for service
	KUBE_SVC_SPEC_TYPE_NODEPORT = "NodePort"
)

// ========================================= kubernetes model - namespace
type KubeNameSpace struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   map[string]string `json:"metadata"` // job module name
}

func (n *KubeNameSpace) SetName(name string) {
	if len(n.Metadata) == 0 {
		n.Metadata = make(map[string]string)
	}
	n.Metadata["name"] = name
}

func (n *KubeNameSpace) GetName() string {
	return n.Metadata["name"]
}

// ========================================= kubernetes model - replicationController
// container
type KubeEnv struct {
	Name  string `json:"name"`  // 环境变量的名称
	Value string `json:"value"` // 环境变量的值
}

type KubePort struct {
	Name          string `json:"name"`          // 端口名称
	ContainerPort int    `json:"containerPort"` // 容器需要监听的端口号
	HostPort      int    `json:"hostPort"`      // 容器所在主机需要监听的端口号，默认与ContainerPort相同
	Protocol      string `json:"protocol"`      // 端口协议，支持TCP&UDP，默认为TCP
}

type KubeVolumeMount struct {
	Name      string `json:"name"`      // 共享存储卷名称
	MountPath string `json:"mountPath"` // 存储卷在容器内Mount的绝对路径
	ReadOnly  bool   `json:"readOnly"`  // 是否为只读模式
}

type KubeResourceLimit struct {
	Cpu    string `json:"cpu"`    // cpu限制条件
	Memory string `json:"memory"` // 内存限制条件
}

type KubeResource struct {
	Limits *KubeResourceLimit `json:"limits"` // 资源限制条件
}

type KubeContainer struct {
	Name            string             `json:"name"`            // 容器名称
	Image           string             `json:"image"`           // 镜像名称 -- job module algorithm 名称
	ImagePullPolicy string             `json:"imagePullPolicy"` // 获取镜像的策略
	Command         map[string]string  `json:"command"`         // 容器的启动命令列表
	WorkDir         string             `json:"workDir"`         // 容器的工作目录
	VolumeMounts    []*KubeVolumeMount `json:"volumeMounts"`    // 可供容器使用的共享存储卷列表
	Ports           []*KubePort        `json:"ports"`           // 容器需要暴露的端口号列表
	Env             []*KubeEnv         `json:"env"`             // 容器运行前需设置的环境变量列表
	Resources       []*KubeResource    `json:"resources"`       // 资源限制条件
}

// volume
type KubePath struct {
	Path string `json:"path"` // pod所在主机的目录
}

type KubeVolume struct {
	Name      string    `json:"name"`      // 共享存储卷名称
	EmptyDidr string    `json:"emptyDidr"` // 默认的存储卷类型，临时目录
	HostPath  *KubePath `json:"hostPath"`  // pod所在主机的目录
}

//
type KubePodSpecNodeSelector struct {
	Key string `json:"key"` // 指定需要调度到的 Node label
}

type KubePodSpecImagePullSecrets struct {
	Name string `json:"name"` // pull镜像时使用的secret名称，以 name=secretkey 格式定义
}

type KubePodSpec struct {
	Containers []*KubeContainer `json:"containers"` // pod中运行的容器列表
	Volumes    []*KubeVolume    `json:"volumes"`    // pod上定义的共享存储卷列表

	RestartPolicy string `json:"restartPolicy"` // pod内容器的重启策略，可选值为Always、OnFailure、Never，默认为Always
	DnsPolicy     string `json:"dnsPolicy"`     // DNS 策略，可选值为 Default，ClusterFirst
	//NodeSelector  PodSpecNodeSelector `json:"nodeSelector"`
	//ImagePullSecrets PodSpecImagePullSecrets `json:"imagePullSecrets"`
}

type KubeSpecSelector struct {
	Name string `json:"name"` //selector，指定pod的管理范围
}

type KubeReplicationControllerSpecTemplate struct {
	MetaData KubeResourceMetaData `json:"metadata"` // 元数据
	Spec     KubePodSpec          `json:"spec"`     // 详细信息
}

type KubeReplicationControllerSpec struct {
	Replicas int                                    `json:"replicas"`
	Selector *KubeSpecSelector                      `json:"selector"` // selector，指定pod的管理范围 -- job module processor 名称
	Template *KubeReplicationControllerSpecTemplate `json:"template"` // 容器的定义
}

type KubeResourceMetaData struct {
	Name        string            `json:"name"`        // 名称
	NameSpace   string            `json:"namespace"`   // 命名空间
	Labels      map[string]string `json:"labels"`      // 自定义标签列表		-- job module processor 名称
	Annotations map[string]string `json:"annotations"` // 自定义注解属性列表
}

type KubeReplicationController struct {
	ApiVersion string                        `json:"apiVersion"`
	Kind       string                        `json:"kind"`
	MetaData   KubeResourceMetaData          `json:"metadata"` //元数据
	Spec       KubeReplicationControllerSpec `json:"spec"`     //详细描述
}

// ========================================= kubernetes model - service
type KubeServiceStatusLoadBalancerIngress struct {
	Ip       string `json:"ip"`       // 外部负载均衡器IP地址
	HostName string `json:"hostname"` // 外部负载均衡器的主机名
}

type KubeServiceStatusLoadBalancer struct {
	Ingress KubeServiceStatusLoadBalancerIngress `json:"ingress"` // 外部负载均衡器
}

type KubeServiceStatus struct {
	LoadBalancer KubeServiceStatusLoadBalancer `json:"loadBalancer"` // 外部负载均衡器
}

type KubeServicePort struct {
	Name       string `json:"name"`       // 端口名称
	Port       int    `json:"port"`       // 服务监听的端口号
	TargetPort int    `json:"targetPort"` // 需要转发到后端Pod的端口号
	NodePort   int    `json:"nodePort"`   // pod所在宿主机的端口号
	Protocol   string `json:"protocol"`   // 端口协议，支持TCP & UDP, 默认为TCP
}

type KubeServiceSpec struct {
	Selector        *KubeSpecSelector  `json:"selector"`        // selector，指定pod的管理范围
	Type            string             `json:"type"`            // service的类型，指定service的访问方式，默认为clusterIp（clusterIp/NodePort/LoadBalancer)
	ClusterIP       string             `json:"clusterIP"`       // 虚拟服务IP地址，type=clusterIp时如果不指定系统自动分配，当type=LoadBalancer时需要指定
	SessionAffinity string             `json:"sessionAffinity"` // 是否支持session，可选值为 ClientIP，默认为空。
	Ports           []*KubeServicePort `json:"ports"`           // service需要暴露的端口号列表
}

//=============================================================================
type KubeService struct {
	ApiVersion string               `json:"apiVersion"`
	Kind       string               `json:"kind"`
	MetaData   KubeResourceMetaData `json:"metadata"` // 元数据
	Spec       KubeServiceSpec      `json:"spec"`     // 详细信息
}

func InitEnv(dbname string, dbdriver string, dbaccount string, dbaddr string, dbconnum int64) error {
	var err error
	var dataSource string

	beego.Debug("init env")

	// register model
	beego.Debug("-->step1 register model")
	orm.RegisterModel(new(User),
		new(Project),
		new(Job),
		new(Module),
		new(Resource),
		new(Pod),
		new(Log),
		new(Algorithm),
		new(Action))

	// register driver
	beego.Debug("-->step2 register driver")
	err = orm.RegisterDriver(dbdriver, orm.DRMySQL)
	if err != nil {
		beego.Debug("register driver", dbdriver, "failed")
		return err
	}
	beego.Debug("success")

	// register default database(mysql)
	beego.Debug("-->step3 register default database(mysql)")
	dataSource = dbaccount + "@" + dbaddr + "/mysql?charset=utf8"
	err = orm.RegisterDataBase("default", dbdriver, dataSource)
	if err != nil {
		beego.Debug("register data base", dataSource, "failed")
		return err
	}
	beego.Debug("success")

	// create database new database
	beego.Debug("-->step4 create new database")
	o := orm.NewOrm()
	_, err = o.Raw("create database if not exists " + dbname + " character set utf8").Exec()
	if err != nil {
		beego.Debug("create database", dbname, "failed!")
		return err
	}
	beego.Debug("success")

	// register new database
	beego.Debug("-->step5 register database new database")
	dataSource = dbaccount + "@" + dbaddr + "/" + dbname + "?charset=utf8"
	err = orm.RegisterDataBase(dbname, dbdriver, dataSource)
	if err != nil {
		beego.Debug("register new database", dbname, "failed!")
		return err
	}
	beego.Debug("success")

	//
	orm.RunCommand()

	//
	orm.RunSyncdb(dbname, false, true)

	return err
}

// func InitEnv() error {
// 	var err error
// 	var dataSource string

// 	beego.Debug("init env")

// 	// register model
// 	beego.Debug("-->step1 register model")
// 	orm.RegisterModel(new(User),
// 		new(Project),
// 		new(Job),
// 		new(Module),
// 		new(Resource),
// 		new(Pod),
// 		new(Log),
// 		new(Algorithm),
// 		new(Action))

// 	// register driver
// 	beego.Debug("-->step2 register driver")
// 	err = orm.RegisterDriver(config.DB_DRIVER_NAME, orm.DRMySQL)
// 	if err != nil {
// 		beego.Debug("register driver", config.DB_DRIVER_NAME, "failed")
// 		return err
// 	}
// 	beego.Debug("success")

// 	// register default database(mysql)
// 	beego.Debug("-->step3 register default database(mysql)")
// 	dataSource = config.DB_ACCOUNT + "@" + config.DB_ACCESS_ADDR + "/mysql?charset=utf8"
// 	err = orm.RegisterDataBase("default", config.DB_DRIVER_NAME, dataSource)
// 	if err != nil {
// 		beego.Debug("register data base", dataSource, "failed")
// 		return err
// 	}
// 	beego.Debug("success")

// 	// create database new database
// 	beego.Debug("-->step4 create new database")
// 	o := orm.NewOrm()
// 	_, err = o.Raw("create database if not exists " + config.DB_NAME + " character set utf8").Exec()
// 	if err != nil {
// 		beego.Debug("create database", config.DB_NAME, "failed!")
// 		return err
// 	}
// 	beego.Debug("success")

// 	// register new database
// 	beego.Debug("-->step5 register database new database")
// 	dataSource = config.DB_ACCOUNT + "@" + config.DB_ACCESS_ADDR + "/" + config.DB_NAME + "?charset=utf8"
// 	err = orm.RegisterDataBase(config.DB_NAME, config.DB_DRIVER_NAME, dataSource)
// 	if err != nil {
// 		beego.Debug("register new database", config.DB_NAME, "failed!")
// 		return err
// 	}
// 	beego.Debug("success")

// 	//
// 	orm.RunCommand()

// 	//
// 	orm.RunSyncdb(config.DB_NAME, false, true)

// 	return err
// }

////============================================================== influxdb 监控数据
type InfluxdbResponseData struct {
	Name    string          `json:"name"`
	Columns []string        `json:"columns"`
	Points  [][]interface{} `json:"points"`
}

type PodResourceUsageValue struct {
	Time  int64  `json:"time"`
	Value string `json:"value"`
}

type PodResourceUsage struct {
	ProjectId   int64                    `json:"projectId"`
	ProjectName string                   `json:"projectName"`
	JobId       int64                    `json:"jobId"`
	JobName     string                   `json:"jobName"`
	ModuleId    int64                    `json:"moduleId"`
	ModuleName  string                   `json:"moduleName"`
	PodId       int64                    `json:"podId"`
	PodName     string                   `json:"podName"`
	Type        string                   `json:"type"`
	Values      []*PodResourceUsageValue `json:"values"`
}

//================================ for new version
type InfluxdbSeries struct {
	Name    string          `json:"name"`
	Columns []string        `json:"columns"`
	Values  [][]interface{} `json:"values"`
}

type InfluxdbResult struct {
	Series []*InfluxdbSeries `json:"series"`
}

type InfluxdbResults struct {
	Results []*InfluxdbResult `json:"results"`
}
