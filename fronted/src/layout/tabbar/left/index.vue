<template>
  <el-icon style="margin-right: 10px">
    <Expand />
  </el-icon>
  <el-breadcrumb separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
    <el-breadcrumb-item>权限管理</el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss"></style>
