<template>
  <div class="home">
    <v-row justify="center">
      <v-dialog v-model="betaDialog" scrollable max-width="600px">
        <v-card>
          <v-card-title>本サイトは，開発途中です</v-card-title>
          <v-card-text>
            不具合・バグなどが存在します．
            正式リリースをお待ちくださいませ．
          </v-card-text>
          <v-card-actions>
            <v-btn color="blue darken-1" text @click="betaDialog = false">閉じる</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog v-model="dialog" scrollable max-width="600px">
        <v-card>
          <v-card-title>質問や回答を行う場合は，ログインが必要です.</v-card-title>
          <v-card-text>
            質問や回答へのいいねを行うにも，ログインが必要になります.<br>
            質問の閲覧やタグ検索などを行うのに，ログインは不要です.
          </v-card-text>
          <v-card-actions>
            <!-- FIXME -->
            <!-- 右にずらしたりをうまくやりたい -->
            <v-row>
              <v-col cols="12" sm="2">
                <v-btn large color="primary" to="/login">ログイン</v-btn>
              </v-col>
              <v-col cols="12" sm="8">
                <v-btn outlined depressed large to="/signup">新規登録</v-btn>
              </v-col>
              <v-col cols="12" sm="2">
                <v-btn color="blue darken-1" text @click="dialog = false">閉じる</v-btn>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-row>
    <div class="main-contents">
      <v-container fluid>
        <v-row>
          <v-col col="12" sm="8">
            <v-row>
              <v-col col="12" sm="10">
                <img src="../assets/rectangle_logo.jpg" alt="ロゴ" width="50%" height="auto">
              </v-col>
            </v-row>
            <div style="margin-left: 1%;">
              <h2>競技プログラミング専用の Q&A サイト</h2>
              <p><router-link class="about" to="./about">このサイトについて</router-link></p>
            </div>
            <QuestionSetting 
              @click="changeMode"
              :curMode="mode"
            />
            <Questions 
              :mode="mode"
            />
          </v-col>
          <v-col col="12" sm="4">
            <SiteInfomation />
            <br>
            <RankingPanel />
          </v-col>
        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';
import QuestionSetting from '@/components/QuestionSetting.vue';
import SiteInfomation from '@/components/SiteInfomation.vue';
import RankingPanel from '@/components/RankingPanel.vue';

@Component({
  components: {
    Questions,
    QuestionSetting,
    SiteInfomation,
    RankingPanel,
  },
})
export default class Home extends Vue {
  private mode: number = 1;  // 新着, 回答数, 閲覧数, いいね, 解決済みかどうか
  private dialog: boolean = false;
  private betaDialog: boolean = false;
  private changeMode(mode: number): void {
    this.mode = mode;
  }
  private created(): void {
    if (this.getToken() === null) {
      this.dialog = true;
    }
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.main-topic {
}
.about {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #0288D1;
  text-decoration: none;
}
.about:hover {
  color: #29B6F6;
  text-decoration: none;
}
</style>
