<template>
  <div class="userpage" v-if="userId">
    <!--{{ userInfo }}-->
    <MyProfile/>
    <v-row>
      <v-col col="12" sm="6">
        <v-tabs
          fixed-tabs
          background-color="transparent"
        >
          <v-tab @click="changeMode('question')">
            <div
              class="display-100 flex-grow-1 text-center"
            >
              <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                質問
              </div>
            </div>
          </v-tab>
          <v-tab @click="changeMode('answer')">
            <div
              class="display-100 flex-grow-1 text-center"
            >
              <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                回答
              </div>
            </div>
          </v-tab>
          <v-tab @click="changeMode('book-mark')">
            <div
              class="display-100 flex-grow-1 text-center"
            >
              <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                ブックマーク
              </div>
            </div>
          </v-tab>
          <v-tab v-if="myId == userId" @click="changeMode('setting')">
            <div
              class="display-100 flex-grow-1 text-center"
            >
              <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                設定
              </div>
            </div>
          </v-tab>
          <v-tab v-else @click="changeMode('setting')" disabled>
            <div
              class="display-100 flex-grow-1 text-center"
            >
              <div style="color: rgb(66, 66, 66);" class="font-weight-bold">
                設定
              </div>
            </div>
          </v-tab>
        </v-tabs>
      </v-col>
    </v-row>
    <v-row>
      <v-col col="12" sm="8">
        <div v-if="mode=='profile'">
        </div>
        <div v-if="mode=='question'">
          <UserQuestions/>
        </div>
        <div v-if="mode=='answer'">
          <UserAnswers/>
        </div>
        <div v-if="mode=='book-mark'">
          <BookMarkedQuestions />
        </div>
        <div v-if="mode=='setting'">
          <MySetting/>
        </div>
      </v-col>
    </v-row>
    <!-- <v-btn @click="getUserInfoByAtCoderUserAPI">get</v-btn> -->
  </div>
  <div v-else>
    <p>ログインしてください．</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import MyProfile from '@/components/MyProfile.vue';
import UserQuestions from '@/components/UserQuestions.vue';
import UserAnswers from '@/components/UserAnswers.vue';
import MySetting from '@/components/MySetting.vue';
import BookMarkedQuestion from '@/components/BookMarkedQuestions';

@Component({
  components: {
    UserAnswers,
    UserQuestions,
    MyProfile,
    MySetting,
    BookMarkedQuestions,
  },
})
export default class UserPage extends Vue {
  private mode: string = 'question';
  private activeBtn: number = 1;

  // 他人
  private userName: string = '';
  private userId: string | Array<(string | null)> = '';

  // 自分
  private myName: string = '';
  private myId: string = '';


  private created(): void {
    if (this.getToken() !== null) {
      const claims = JSON.parse(atob(this.getToken().split('.')[1]));
      this.myName = claims.name;
      this.myId = claims.uid;
    }
    this.userId = this.$route.query.uid;

    const url = '/api/no-auth/user/' + String(this.userId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
    });
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
  private changeMode(mode: string): void {
    this.mode = mode;
  }
}
</script>
