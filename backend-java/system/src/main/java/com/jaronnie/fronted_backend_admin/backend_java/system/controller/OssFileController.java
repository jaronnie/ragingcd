package com.jaronnie.fronted_backend_admin.backend_java.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.OssFileVo;
import com.jaronnie.fronted_backend_admin.backend_java.common.util.R;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.RequiredArgsConstructor;
import org.dromara.x.file.storage.core.FileInfo;
import org.dromara.x.file.storage.core.FileStorageService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.util.Date;

@Api(tags = "oss 文件管理")
@RestController
@RequiredArgsConstructor
@RequestMapping("/api/v1.0/system/oss_file")
public class OssFileController {
    private final FileStorageService fileStorageService;//注入实列

    @ApiOperation(value = "上传文件")
    @SaCheckLogin
    @PostMapping("/upload")
    public R<OssFileVo> upload(MultipartFile file) {
        FileInfo fileInfo = fileStorageService.of(file).upload();
        return R.ok(OssFileVo.builder()
                        .filename(fileInfo.getFilename())
                        .url(fileInfo.getUrl())
                        .uploadTime(new Date())
                        .build());
    }
}
