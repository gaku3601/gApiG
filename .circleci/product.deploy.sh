docker login registry.gitlab.com -u ${UserName} -p ${Password}

# コンテナのUPLOAD
echo $CIRCLE_TAG
ProjectName="gapig"
ContainerName="management-api"
docker build -t registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG} ./management-api/.
docker tag registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG} registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:latest
docker push registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG}
docker push registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:latest

ContainerName="backendapp"
docker build -t registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG} ./backendapp/.
docker tag registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG} registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:latest
docker push registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:${CIRCLE_TAG}
docker push registry.gitlab.com/gaku3601/${ProjectName}/${ContainerName}:latest
