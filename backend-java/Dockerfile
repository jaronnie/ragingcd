FROM openjdk:8-jre-alpine

WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 8082

ADD ./target/*.jar ./app.jar

RUN echo "profile: ${profile}"
ENTRYPOINT ["sh","-c","java ${JAVA_OPTS} -Dspring.profiles.active=${profile} -Djava.security.egd=file:/dev/./urandom -jar app.jar"]