<template>
  <div class="myprofile">
    <v-card>
      <v-tabs vertical height="400" hide-slider>
        <v-tab>
          <v-icon left>mdi-account</v-icon>
          基本設定
        </v-tab>
        <v-tab>
          <v-icon left>mdi-delete</v-icon>
          アカウント削除
        </v-tab>
        <!-- v-tab>
          <v-icon left>mdi-access-point</v-icon>
          Option 3
        </v-tab-->

        <v-tab-item>
          <v-card flat>
            <UserUpdateForms />
          </v-card>
        </v-tab-item>
        <v-tab-item>
          <v-card flat>
            <WithdrawButton />
          </v-card>
        </v-tab-item>
        <v-tab-item>
          <v-card flat>
            <v-card-text>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-card>

  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import WithdrawButton from '@/components/WithdrawButton.vue';
import UserUpdateForms from '@/components/UserUpdateForms.vue';

@Component({
  components: {
    WithdrawButton,
    UserUpdateForms,
  },
})

export default class MySetting extends Vue {
}
</script>
