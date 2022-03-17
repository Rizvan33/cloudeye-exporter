package collector

import (
	"encoding/json"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"

	"github.com/huaweicloud/cloudeye-exporter/logs"
)

type GaussdbNodeInfo struct {
	ResourceBaseInfo
	NodeProperties
}

type NodeProperties struct {
	InnerPort  string                   `json:"innerPort"`
	InnerIp    string                   `json:"innerIp"`
	Role       string                   `json:"role"`
	EngineName string                   `json:"engineName"`
	Dimensions []model.MetricsDimension `json:"dimensions"`
}

var gaussdbInfo serversInfo

func (exporter *BaseHuaweiCloudExporter) getGaussdbResourceInfo() (map[string]labelInfo, []model.MetricInfoList) {
	resourceInfos := map[string]labelInfo{}
	filterMetrics := make([]model.MetricInfoList, 0)
	gaussdbInfo.Lock()
	defer gaussdbInfo.Unlock()
	if gaussdbInfo.LabelInfo == nil || time.Now().Unix() > gaussdbInfo.TTL {
		sysConfigMap := getMetricConfigMap("SYS.GAUSSDB")
		if metricNames, ok := sysConfigMap["gaussdb_mysql_instance_id,gaussdb_mysql_node_id"]; ok {
			nodes, err := getAllGaussdbNodesFromRMS()
			if err == nil {
				for _, node := range nodes {
					metrics := buildDimensionMetrics(metricNames, "SYS.GAUSSDB", node.Dimensions)
					filterMetrics = append(filterMetrics, metrics...)
					info := labelInfo{
						Name:  []string{"name", "epId", "innerPort", "innerIp", "role", "engineName"},
						Value: []string{node.Name, node.EpId, node.InnerPort, node.InnerIp, node.Role, node.EngineName},
					}
					keys, values := getTags(node.Tags)
					info.Name = append(info.Name, keys...)
					info.Value = append(info.Value, values...)
					resourceInfos[GetResourceKeyFromMetricInfo(metrics[0])] = info
				}
			}
		}

		gaussdbInfo.LabelInfo = resourceInfos
		gaussdbInfo.FilterMetrics = filterMetrics
		gaussdbInfo.TTL = time.Now().Add(TTL).Unix()
	}
	return gaussdbInfo.LabelInfo, gaussdbInfo.FilterMetrics
}

func getAllGaussdbNodesFromRMS() ([]GaussdbNodeInfo, error) {
	resp, err := listResources("gaussdb", "nodes")
	if err != nil {
		logs.Logger.Errorf("Failed to list resource of gaussdb.nodes, error: %s", err.Error())
		return nil, err
	}
	nodes := make([]GaussdbNodeInfo, 0, len(resp))
	for _, resource := range resp {
		properties, err := fmtGaussdbNodeProperties(resource.Properties)
		if err != nil {
			continue
		}
		nodes = append(nodes, GaussdbNodeInfo{
			ResourceBaseInfo: ResourceBaseInfo{
				ID:   *resource.Id,
				Name: *resource.Name,
				EpId: *resource.EpId,
				Tags: resource.Tags,
			},
			NodeProperties: *properties,
		})
	}
	return nodes, nil
}

func fmtGaussdbNodeProperties(properties map[string]interface{}) (*NodeProperties, error) {
	bytes, err := json.Marshal(properties)
	if err != nil {
		logs.Logger.Errorf("Marshal gaussdb node properties error: %s", err.Error())
		return nil, err
	}
	var nodeproperties NodeProperties
	err = json.Unmarshal(bytes, &nodeproperties)
	if err != nil {
		logs.Logger.Errorf("Unmarshal to NodeProperties error: %s", err.Error())
		return nil, err
	}

	return &nodeproperties, nil
}
