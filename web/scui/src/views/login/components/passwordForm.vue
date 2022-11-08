<template>
  <el-form
    ref="loginForm"
    :model="form"
    :rules="rules"
    label-width="0"
    size="large"
  >
    <el-form-item prop="user">
      <el-input
        v-model="form.user"
        prefix-icon="el-icon-user"
        clearable
        :placeholder="$t('login.userPlaceholder')"
      >
        <!-- <template #append>
					<el-select v-model="userType" style="width: 130px;">
						<el-option :label="$t('login.admin')" value="admin"></el-option>
						<el-option :label="$t('login.user')" value="user"></el-option>
					</el-select>
				</template> -->
      </el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        v-model="form.password"
        prefix-icon="el-icon-lock"
        clearable
        show-password
        :placeholder="$t('login.PWPlaceholder')"
      ></el-input>
    </el-form-item>

    <el-form-item prop="captchaValue">
      <el-col :span="14">
        <el-input
          v-model="form.captchaValue"
          placeholder="验证码"
          prefix-icon="el-icon-ChatLineSquare"
        ></el-input>
      </el-col>
      <el-col :span="6" class="captchaImg">
        <el-image :src="form.captchaUrl"></el-image>
      </el-col>
      <el-col :span="2"
        ><el-button @click="referCaptcha" text>刷新</el-button></el-col
      >
    </el-form-item>
    <el-form-item style="margin-bottom: 10px">
      <el-col :span="12">
        <el-checkbox
          :label="$t('login.rememberMe')"
          v-model="form.autologin"
        ></el-checkbox>
      </el-col>
      <el-col :span="12" class="login-forgot">
        <!-- <router-link to="/reset_password">{{ $t('login.forgetPassword') }}？</router-link> -->
      </el-col>
    </el-form-item>
    <el-form-item>
      <el-button
        type="primary"
        style="width: 100%"
        :loading="islogin"
        round
        @click="login"
        >{{ $t("login.signIn") }}</el-button
      >
    </el-form-item>
    <div class="login-reg">
      <!-- {{$t('login.noAccount')}} <router-link to="/user_register">{{$t('login.createAccount')}}</router-link> -->
    </div>
  </el-form>
</template>

<script>
export default {
  data() {
    return {
      userType: "admin",
      form: {
        user: "admin",
        password: "admin",
        autologin: false,
        captchaId: "",
        captchaValue: "",
        captchaUrl: "",
      },
      rules: {
        user: [
          {
            required: true,
            message: this.$t("login.userError"),
            trigger: "blur",
          },
        ],
        captchaValue: [
          {
            required: true,
            message: "验证码必须",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: this.$t("login.PWError"),
            trigger: "blur",
          },
        ],
      },
      islogin: false,
    };
  },
  watch: {},
  mounted() {
    this.captchaInfo();
  },
  methods: {
    referCaptcha() {
      this.captchaInfo();
    },
    async captchaInfo() {
      let res = await this.$API.common.captchaInfo.get();
      if (res.code === 200) {
        this.form.captchaId = res.data.id;
        this.form.captchaUrl = res.data.url + "/100/40";
      }
    },
    async login() {
      var validate = await this.$refs.loginForm.validate().catch(() => {});
      if (!validate) {
        return false;
      }

      this.islogin = true;
      var data = {
        username: this.form.user,
        password: this.form.password,
        captchaValue: this.form.captchaValue,
        captchaId: this.form.captchaId,
      };
      //获取token
      var user = await this.$API.auth.token.post(data);
      if (user.code == 200) {
        this.$TOOL.cookie.set("TOKEN", user.data.token, {
          expires: this.form.autologin ? 24 * 60 * 60 : 0,
        });
        this.$TOOL.data.set("USER_INFO", user.data.userInfo);
      } else {
							 this.referCaptcha();
        this.islogin = false;
        this.$message.warning(user.message);
        return false;
      }
      //获取菜单
      var menu = await this.$API.system.menu.myMenus.get();
      if (menu.code == 200) {
        if (menu.data.menu.length == 0) {
          this.islogin = false;
          this.$alert(
            "当前用户无任何菜单权限，请联系系统管理员",
            "无权限访问",
            {
              type: "error",
              center: true,
            }
          );
          return false;
        }
        this.$TOOL.data.set("MENU", menu.data.menu);
        this.$TOOL.data.set("PERMISSIONS", menu.data.permissions);
      } else {
        this.islogin = false;
        this.$message.warning(menu.message);
        return false;
      }

      this.$router.replace({
        path: "/",
      });
      this.$message.success("Login Success 登录成功");
      this.islogin = false;
    },
  },
};
</script>

<style scoped>
.captchaImg {
  display: flex;
  flex-direction: row;
  align-content: center;
}
</style>
