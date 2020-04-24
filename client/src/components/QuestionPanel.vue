<template>
  <div class="question-panel">
    <v-card
      class="mx-auto"
      color="white"
    >
      <div v-if="isReady">
        <v-card-title>
          <router-link class="title" :to="{ name: 'question', query: { questionId: this.questionId }}" >{{ question.title }}</router-link>
        <v-spacer />
        <TwitterIcon
          :twitterId="userTwitterId"
          size="36"
          apiSize="n"
        />
        <span style="font-size: 14px;">
          <UserName
            :name="userName"
            :uid="question.uid"
          />
        </span>
        </v-card-title>
        <!--v-btn
          class="tag"
          color="blue-grey lighten-4"
          x-small
        >
          タグ {{ question.tid }}
        </v-btn-->
        <v-divider class="mx-4"></v-divider>
        <v-row style="margin-left: 1%;">
          <v-chip
            outlined
            class="ma-2"
            color="rgb(66, 66, 66, 0.6)"
            text-color="rgb(66, 66, 66, 0.6)"
          >
            <span class="chip-text">
              回答数{{ question.answerCount }}
            </span>
          </v-chip>
          <v-chip
            outlined
            class="ma-2"
            color="rgb(66, 66, 66, 0.6)"
            text-color="rgb(66, 66, 66, 0.6)"
          >
            <span class="chip-text">
              閲覧数{{ question.browseCount }}
            </span>
          </v-chip>
          <v-chip
            outlined
            class="ma-2"
            color="rgb(66, 66, 66, 0.6)"
            text-color="rgb(66, 66, 66, 0.6)"
          >
            <span class="chip-text">
              いいね{{ question.favoriteCount }}
            </span>
          </v-chip>
          <v-chip
            v-if="question.completed"
            outlined
            class="ma-2"
            color="rgb(116, 181, 103)"
            text-color="rgb(116, 181, 103)"
          >
            <span class="chip-text">
              解決済み
            </span>
          </v-chip>
          <v-chip
            v-else
            outlined
            class="ma-2"
            color="rgb(231,175,95)"
            text-color="rgb(231,175,95)"
          >
            <span class="chip-text">
              未解決
            </span>
          </v-chip>

          <!--SquarePanel 
            message="回答数"
            :num="question.answerCount"
            color="blue-grey"
          />
          <SquarePanel
            message="閲覧数"
            :num="question.browseCount"
            color="blue-grey"
          />
          <SquarePanel
            message="いいね"
            :num="question.favoriteCount"
            color="blue-grey"
          />
          <SquarePanel
            v-if="question.completed"
            message="解決済"
            color="#5cb85c"
          />
          <SquarePanel
            v-if="!question.completed"
            message="未解決"
            color="#f0ad4e"
          /-->
          <div class="flex-grow-1"></div>
          <div style="margin-left:auto; margin-top: 20px; margin-right: 20px;">
            <span class="date">{{ question.date }}</span>
          </div>
        </v-row>
      </div>
      <div v-else class="text-center">
        <v-progress-circular
          :size="100"
          color="primary"
          indeterminate
        ></v-progress-circular>
      </div>
    </v-card>
    
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit, Watch } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';
import UserName from '@/components/UserName.vue';
import TwitterIcon from '@/components/TwitterIcon.vue';
import { Question } from '@/models/Question.ts';

@Component({
  components: {
    UserName,
    SquarePanel,
    TwitterIcon,
  },
})
export default class QuestionPanel extends Vue {
  @Prop({required: true})
  private questionId!: number;

  // データベース Question 通り
  private question: Question = {
    id: 0,
    title: 'title title title title title',
    body: 'body body body body body',
    completed: false,
    date: '',

    favoriteCount: 0,
    answerCount: 0,
    browseCount: 0,
  };

  private isReady: boolean = false;

  // 質問者の名前
  private userName: string = '';

  // 質問者の twitter Id
  private userTwitterId: string = '';

  // コンポーネントが作られた時に呼び出される関数
  private created(): void {
    this.init();
  }

  // questionId が変更されたらば,表示を変える
  @Watch('questionId')
  private idChanged(): void {
    this.init();
  }

  private init(): void {
    // TODO api/question/:id (GET) を叩く
    const url = '/api/no-auth/question/' + String(this.questionId);
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.question = json;
      this.setUser();
    });
  }

  private setUser(): void {
    const url = 'api/no-auth/user/' + String(this.question.uid);
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
      this.userTwitterId = json.twitter_id;
      this.isReady = true;
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }

}
</script>

<style scoped>
.question-panel {
  margin: 3%;
}
.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}
.title {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #0288D1;
  text-decoration: none;
}
.title:hover {
  color: #29B6F6;
  text-decoration: none;
}
.small {
  font-size: 75%;
}
.date {
  color: rgb(151, 151, 151);
}
.chip-text {
  font-weight: bold;
}
</style>
