# Go with ElasticSearch

Application created to work with ElasticSearch engine

### Based on articles
* [How To Create An Elasticsearch Index Using The Olivere Driver In Golang](https://kb.objectrocket.com/elasticsearch/how-to-create-an-elasticsearch-index-using-the-olivere-driver-in-golang-548#an+example+of+a+golang+olivere+driver+and+elasticsearch+connection)
* [Creating Elastic Search Mapping in Golang](https://medium.com/terragoneng/creating-elastic-search-mapping-in-golang-654f221c4e4b)
* [How to implement Elasticsearch in Go](https://www.freecodecamp.org/news/go-elasticsearch/)

### Prerequisites
* ElasticSearch server running on `localhost:9200`
```
# Download docker image
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.12.0

# Run ElasticSearch node
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.12.0
```

### How to run in CMD
1. Prepare prerequisite
2. go build go-elastic
3. go run go-elastic
4. Use postman collections as template