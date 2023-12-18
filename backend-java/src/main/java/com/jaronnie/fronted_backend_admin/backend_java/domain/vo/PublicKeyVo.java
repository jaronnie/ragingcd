package com.jaronnie.fronted_backend_admin.backend_java.domain.vo;

import io.swagger.annotations.ApiModelProperty;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class PublicKeyVo {
    @ApiModelProperty("秘钥类型 RSA or SM2")
    private String type;
    @ApiModelProperty("公钥-PKCS8")
    private String publicKey;
}
