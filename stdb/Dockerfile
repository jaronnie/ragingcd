FROM openjdk:8-jre-alpine

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 8082

ADD ./admin/target/*.jar ./app.jar

ENTRYPOINT ["sh","-c","java ${JAVA_OPTS} -Dspring.profiles.active=docker -Djava.security.egd=file:/dev/./urandom -jar app.jar"]