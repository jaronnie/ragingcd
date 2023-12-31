package com.jaronnie.ragingcd.stdb.system.controller;

import cn.dev33.satoken.annotation.SaCheckLogin;
import com.jaronnie.ragingcd.stdb.system.domain.vo.OssFileVo;
import com.jaronnie.ragingcd.stdb.common.util.R;
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
public class SystemOssFileController {
    private final FileStorageService fileStorageService;

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
