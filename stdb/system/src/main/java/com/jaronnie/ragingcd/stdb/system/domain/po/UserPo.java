package com.jaronnie.ragingcd.stdb.system.domain.po;

import com.baomidou.mybatisplus.annotation.FieldStrategy;
import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.*;

import javax.persistence.Entity;
import javax.persistence.Id;
import java.io.Serializable;

@EqualsAndHashCode(callSuper = false)
@TableName(value = "user", autoResultMap = false)
@Entity
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class UserPo extends BaseEntity implements Serializable {
    private static final long serialVersionUID = 1L;

    @Id
    @TableId
    private Integer id;

    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String username;
    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String password;
    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String avatar;
    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String email;
}
