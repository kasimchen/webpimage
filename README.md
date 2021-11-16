1、根据传输过去的json图片链接 获取到图片的大小和类型
2、webp图片转png图片
3、视频转发

docker-composer up -d





```shell
curl --location --request GET 'http://127.0.0.1:8000/webp_to_image?url=https://ci.xiaohongshu.com/1492f4aa-1848-7707-4362-850c8c3d9891?imageView2/2/w/100/h/100/q/75/format/webp' \
--form 'json="[\"https://ci.xiaohongshu.com/e6e3f1ff-d752-716f-8c40-f98618e5f83d?imageView2/2/w/200/h/200/q/75/format/webp\",\"https://img.alicdn.com/imgextra/i4/725677994/O1CN01ZOtZcI28vIgqsGv39_!!725677994.jpg_430x430q90.jpg\"]"'

```


```
curl --location --request POST 'http://127.0.0.1:8000/get_image_infos?json=11' \
--form 'json="[\"https://hbfile.huabanimg.com/img/category_page/395e6de934a06825e38e608bae05f6e4f81f5501e186\",\"https://hbfile.huabanimg.com/img/category_page/58e3aca8638426231c9c1624c98636c95aa3a063d9bc\",\"https://hbfile.huabanimg.com/img/category_page/6a9dadf20abddc3701e44905876b0951fdeea7bdbe5e\",\"https://hbfile.huabanimg.com/img/category_page/3bcf35f7b6445c48ca3e00c5fbc28983973d030aaaac\",\"https://hbfile.huabanimg.com/img/category_page/ab6755376ce1408465462b027ec1381cdc8d7368bbe8\",\"https://hbfile.huabanimg.com/img/category_page/f92bff7c9fc23a955cd23ab0de38ed2af0ff9ce78864\",\"https:/img/quick_pin.svg\",\"https://hbimg.huabanimg.com/72c91e3c0158874d0ce7b4ef601512fe648c66fb1a1b7-4qA1e3_fw236/format/webp\",\"https://hbimg.huabanimg.com/db6b969b3dc5ab570c07db054f975febde7b9add1becc-m7XGqX_sq75/format/webp\",\"https://hbimg.huabanimg.com/654953460733026a7ef6e101404055627ad51784a95c-B6OFs4_sq75/format/webp\",\"https://hbimg.huabanimg.com/9a210665e62c4f06acc3bbda189763982d3b79e8dc8c-jBiakj_fw236/format/webp\",\"https://hbimg.huabanimg.com/871b290e7738b5a31fc463cb3f96e5d7849d50714031-nSw9zE_sq75/format/webp\",\"https://hbimg.huabanimg.com/c67868b33fff328851174ee5a82beed09ebc81f05f8abd-KkAK43_fw236/format/webp\",\"https://hbimg.huabanimg.com/5ce683c3ee73c04dcdc6b0a2753d8714cec176f75e71e-hIIvyB_sq75/format/webp\",\"https://hbimg.huabanimg.com/1bab1587ce79e7e0ebb51a03780ba19e61309163a1377-nnPJIQ_fw236/format/webp\",\"https://hbimg.huabanimg.com/badcac1ed1c05c6c7ca275b164c4522fadeaf4c59b6-X1GXAD_sq75/format/webp\",\"https://hbimg.huabanimg.com/a3030c3f038e5e9449867fbfb7ddda982a14e02142993-488i27_fw236/format/webp\",\"https://hbimg.huabanimg.com/bbee3f965f35b122ca154719b6bb416b5d98251cbf1-qt40cI_sq75/format/webp\",\"https://hbimg.huabanimg.com/85b44124d9a8f979ac5cc70a3037701ebb56e6cf1bf5e3-RgbMxU_fw236/format/webp\",\"https://hbimg.huabanimg.com/e106b82cc413305b92ee0bf0c2621d78593d295dfd3-vqeTSE_sq75/format/webp\",\"https://hbimg.huabanimg.com/7a705df2273680540e11060fd52582460a19727111bad2-3ByDyq_fw236/format/webp\",\"https://hbimg.huabanimg.com/4c904cd3be38673abcedca1553db5fe55a54a73f6a8a-CrqIro_sq75/format/webp\",\"https://hbimg.huabanimg.com/a22e2e733121713035104f65df2a55294843e490a5c2e-ZHrjgt_fw236/format/webp\",\"https://hbimg.huabanimg.com/5851c354faebcd44186728b2792fc7132dd0dc1cbfe70-3lTETx_sq75/format/webp\",\"https://hbimg.huabanimg.com/afab8605323e5ff7554a735e2f21dd433e0d8d7754a6-diUBdS_fw236/format/webp\",\"https://hbimg.huabanimg.com/c804746d882cf06e667abc4fddb75e89d4c25685cc7e-qRMOVj_sq75/format/webp\",\"https://hbimg.huabanimg.com/44f0c07e9bc84f004284b2c00e507ef0d406c20c5feb6-cjksBz_fw236/format/webp\",\"https://hbimg.huabanimg.com/850f84574a90fbcc0aeffae8c99c9da6fb06ebd2123d-jQ9RQB_sq75/format/webp\",\"https://hbimg.huabanimg.com/d3d959f1aabdc6e38c757fa501a69abf7f7d2d33c939b-QKXVZv_fw236/format/webp\",\"https://hbimg.huabanimg.com/123ec0d3833f513ac56d5b6fcabae40432a58f16174-UYQtq4_sq75/format/webp\",\"https://hbimg.huabanimg.com/8d683a2a88636881a6f8b591a29b91c5f4956f9c80a36-JWqByI_fw236/format/webp\",\"https://hbimg.huabanimg.com/0d1dcc7890610682db0b902e6aeb13908302b0a684436-QEsT9R_sq75/format/webp\",\"https://hbimg.huabanimg.com/e08aff170ead59463b3263b958e2929301db9e0d3fef-W3r10x_fw236/format/webp\",\"https://hbimg.huabanimg.com/f8fb5b0c1b3a877ad8e80aec4a808a5073d57d122198a-VIart0_sq75/format/webp\",\"https://hbimg.huabanimg.com/71b6d47436dcd390f41793514414c01b607cf2c2f587-QKhivX_fw236/format/webp\",\"https://hbimg.huabanimg.com/597c40a5f23638290c696845aafcf254928222a9af6-i4zsMn_sq75/format/webp\",\"https://hbimg.huabanimg.com/eb92c148397fa4b3ba86930c9ebc52b3e92372421de41-2bgXno_fw236/format/webp\",\"https://hbimg.huabanimg.com/6d78203cc43702101a6cee810e0f2cccd336edf511b11-eJ571U_sq75/format/webp\",\"https://hbimg.huabanimg.com/1f9439c4b8d6d6bbbdea55b35fda228c725d65a0bdcf-WxiOpB_fw236/format/webp\",\"https://hbimg.huabanimg.com/5b971bfccf65e734efe4b83c489e6319767beea85326c-H4U6rp_sq75/format/webp\",\"https://hbimg.huabanimg.com/a5f4f2a48062b51131b816664c10a168d9d69c061c1e3-H7ninN_fw236/format/webp\",\"https://hbimg.huabanimg.com/5f9a46d0863bb9bb11260ec7756903417936364ac191-vdKsdT_sq75/format/webp\",\"https://hbimg.huabanimg.com/533067d3caf76f2908f420eaa0e60df5e7d2e2038f6e7-DMIVeI_fw236/format/webp\",\"https://hbimg.huabanimg.com/69009c4b49c823aaa8d1d404e11db8d827d157dbef00-qK6RFk_sq75/format/webp\",\"https://hbimg.huabanimg.com/205e243a76ab62981d6958214a945cc273cab5a696e37-TbTgwG_fw236/format/webp\",\"https://hbimg.huabanimg.com/e962bec3d6edad5d5a5c3e515e323fac87393bdb939-SseFDa_sq75/format/webp\",\"https://hbimg.huabanimg.com/37f15ed801bf583924a58cd09ba7834ccc2de42f7d48e-fqfdfC_fw236/format/webp\",\"https://hbimg.huabanimg.com/c54a32721eb0d3aef19a599bace3b6cbb0b9e6a250ef1-yN1iVd_sq75/format/webp\",\"https:/img/loading.gif\",\"https://hbimg.huabanimg.com/9a210665e62c4f06acc3bbda189763982d3b79e8dc8c-jBiakj_fw658/format/webp\",\"https://hbimg.huabanimg.com/9a210665e62c4f06acc3bbda189763982d3b79e8dc8c-jBiakj_fw86/format/webp\",\"https://hbimg.huabanimg.com/4a927d2f5a2c9fb5dc3231d497626947b9abedb514e8f-Oov5DE_fw86/format/webp\",\"https://hbimg.huabanimg.com/cff7555f7e11323936996b50441ec1388e2dbd742110c-0OhahS_fw86/format/webp\",\"https://hbimg.huabanimg.com/a8cc8d2d21b481781191e71c378f05fdd38149272fd61-mujYvt_fw86/format/webp\",\"https://hbimg.huabanimg.com/fdca0b42e331136509b4e2547cb26842c8d681dc744c8-19evmS_fw86/format/webp\",\"https://hbimg.huabanimg.com/e6034092cc63bd078ad47916894658abe46a69f723c02-W1raFF_fw86/format/webp\",\"https://hbimg.huabanimg.com/1d310c56a2b35bedbb670007ae19287d3a8633af1823c-CroKW8_fw86/format/webp\",\"https://hbimg.huabanimg.com/47fdea8e5123e522b44f771cea288ecf60bff76714b84-v4yZWw_fw86/format/webp\",\"https://hbimg.huabanimg.com/858720f2555a7b41682e47280042eb50c5ea5ca652d4b-fgzBDv_fw86/format/webp\",\"https://hbimg.huabanimg.com/af479b2878585725de1381420576a7b2bf39b4551d2e5-RQh5rO_fw86/format/webp\",\"https://hbimg.huabanimg.com/d3dff2f6458f1f7b274e1adf07cba57b08e50c4d1b48c-AX27rN_fw86/format/webp\",\"https://hbimg.huabanimg.com/dc079e6e43c780a899c3526b70f6c2c5a73f5a8a7e3f-ZZ9feh_fw86/format/webp\",\"https://hbimg.huabanimg.com/e925d448274db22140d2ce66fa7c832e1039a133f34c-iNy6hc_fw86/format/webp\",\"https://hbimg.huabanimg.com/c5ae2a944b6004fc0cba13ac0210f79cdd3f6c47fff7-I9Ll5c_fw86/format/webp\",\"https://hbimg.huabanimg.com/a57e6367ffa72224e27428a9a08d36d822ee558e5b9b1-x5V6al_fw86/format/webp\",\"https://hbimg.huabanimg.com/dac5ee8227be66f75b1cce239b1eef1b66e66e5541148-rXHHd7_fw86/format/webp\",\"https://hbimg.huabanimg.com/c337ee16c60bf63d33769ec241a9ef34a729612080a94-foHSDZ_fw86/format/webp\",\"https://hbimg.huabanimg.com/d423ccf6460ab57595693fb75fbca2228bb8e4861297e-XY7wXr_fw86/format/webp\",\"https://hbimg.huabanimg.com/b04a99e8103bb2827f51abb6f0fe2676b3a9b10d18aa-qx2c7O_fw86/format/webp\",\"https://hbimg.huabanimg.com/cbf76311cffcfe4bb87868bb31779ae622774c25204e0-JwkKKe_fw86/format/webp\",\"https://d5nxst8fruw4z.cloudfront.net/atrk.gif?account=7I54i1acVE00Uu\",\"https://static-ud.udesk.cn/img/msg2@68x66.png\",\"https://hbfile.huabanimg.com/img/activity/f5848f389a383c60410e76323f9ae406896723b99cbc1!/fw/570\",\"https://hbfile.huabanimg.com/img/activity/6c215bc617877f4a5a4d6f3f9753b5616fa8a9551cab34!/fw/570\",\"https://hbimg-other.huabanimg.com/ca6643c04c7fe05d63322327e91f6ace189d4a391636093587VJjrNyGr\"]"'
```



http://127.0.0.1:8000/proxy?url=转发的视频地址
