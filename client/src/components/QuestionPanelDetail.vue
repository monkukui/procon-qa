<template>
  <div class="question-panel-detail">
    <v-card
      class="mx-auto"
    >
      <div v-if="isReady">
        <v-list-item>
          <v-list-item-content>
            <div class="question-title">
              {{ question.title }}
            </div>
            <v-row>
              <v-col cols="12" sm="3">
                <span v-if="question.completed">
                  <v-chip
                    label
                    class="ma-2"
                    color="rgb(116, 181, 103)"
                    text-color="white"
                  >
                    解決済みの質問
                  </v-chip>
                </span>
                <span v-else>
                  <v-chip
                    label
                    class="ma-2"
                    color="rgb(231, 175, 95)"
                    text-color="white"
                  >
                    未解決の質問
                  </v-chip>
                </span>
              </v-col>
            </v-row>

          </v-list-item-content>
        </v-list-item>
        <v-divider class="mx-4"></v-divider>
        <v-row>
          <v-col cols="12" sm="12">
            <v-card-text>
              <div class="mavon-editor">
                <mavon-editor 
                  v-model="question.body"
                  defaultOpen="preview"
                  :toolbarsFlag="fa"
                  :subfield="fa"
                  :boxShadow="fa"
                />
              </div>
            </v-card-text>
          </v-col>
          <v-col v-if="question.url" cols="12" sm="12">
            <div class="url">
              <span v-if="question.url">
                URL: <a :href="question.url">{{ question.url }}</a>
              </span>
            </div>
          </v-col>
        </v-row>
        <div v-if="name == userName">
          <v-btn 
            v-if="question.completed"
            @click="updateQuestionCompleted" 
            class="ma-2"
            tile
            outlined 
            color="rgb(223, 177, 109)"
            rounded
          >
            <v-icon>mdi-undo-variant</v-icon>未解決に戻す
          </v-btn>
          <v-btn 
            v-else
            @click="updateQuestionCompleted" 
            class="ma-2"
            tile
            outlined 
            color="rgb(131, 179, 112)"
            rounded
          >
            <v-icon>mdi-check</v-icon>解決済みにする
          </v-btn>
        </div>
        <v-card-actions>
          <TwitterIcon
            :twitterId="userTwitterId"
            :uid="question.uid"
            size="60"
            apiSize="b"
          />
          <v-card-text>
            投稿者: 
            <span style="font-size: 14px;">
              <UserName
                :name="userName"
                :uid="question.uid"
              />
            </span>
            <br>
            投稿日時: {{ question.date }}
          </v-card-text>
        </v-btn>
        <v-menu left :offset-x="tr">
          <template v-slot:activator="{ on }">
            <v-btn icon 
              color="error"
              :disabled="name != userName"
              x-large
              v-on="on"
            >
              <v-icon>mdi-delete</v-icon>
            </v-btn> 
          </template>
          <v-card>
            <v-list>
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>本当に削除しますか？</v-list-item-title>
                  <v-list-item-subtitle>この操作は取り消せません．</v-list-item-subtitle>
                </v-list-item-content>
              </v-list-item>
            </v-list>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="menu = false">キャンセル</v-btn>
              <v-btn color="error" @click="deleteQuestion">削除する</v-btn>
            </v-card-actions>
          </v-card>
        </v-menu>
          <!--v-tooltip top>
            <template v-slot:activator="{ on }">
              <v-btn icon 
                color="error"
                :disabled="name != userName"
                x-large
                v-on="on"
              >
                <v-icon @click="alert = !alert">mdi-delete</v-icon>
              </v-btn> 
            </template>
            <span>消去</span>
          </v-tooltip-->

          <v-tooltip top>
            <template v-slot:activator="{ on }">
              <v-btn 
                v-if="isFavorited"
                icon
                x-large
                color="pink"
                :disabled="userName === name || userName === '' || name == ''"
                v-on="on"
              >
                <v-icon @click="favoriteQuestion">mdi-heart</v-icon>
                <span v-if="userName !== ''">
                  {{ question.favoriteCount }}
                </span>
              </v-btn>
              <v-btn 
                v-else
                icon
                x-large
                color="pink lighten-4"
                :disabled="userName === name || userName === '' || name == ''"
                v-on="on"
              >
                <v-icon @click="favoriteQuestion">mdi-heart</v-icon>
                <span v-if="userName !== ''">
                  {{ question.favoriteCount }}
                </span>
              </v-btn>
            </template>
            <span>いいね</span>
          </v-tooltip>
          <v-tooltip top>
            <template v-slot:activator="{ on }">
              <v-btn 
                v-if="isBookMarked"
                icon
                x-large
                color="blue"
                :disabled="name == ''"
                v-on="on"
              >
                <v-icon @click="bookMark">mdi-bookmark</v-icon>
              </v-btn>
              <v-btn 
                v-else
                icon
                x-large
                color="blue lighten-4"
                :disabled="name == ''"
                v-on="on"
              >
                <v-icon @click="bookMark">mdi-bookmark</v-icon>
              </v-btn>
            </template>
            <span>ブックマーク</span>
          </v-tooltip>
        </v-card-actions>
        <!--v-row>
          <v-col md="12">
          </v-col>
        </v-row-->
      </div>
      <div v-else class="text-center">
        <v-progress-circular
          :size="100"
          color="primary"
          indeterminate
        ></v-progress-circular>
      </div>

      <!-- コメント -->
      <v-col
        md="10"
        class="ml-auto"
      >
        <QuestionComments 
          :qid="questionId"
          :uid="question.uid"
        />
      </v-col>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Emit } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';
import UserName from '@/components/UserName.vue';
import TwitterIcon from '@/components/TwitterIcon.vue';
import QuestionComments from '@/components/QuestionComments.vue';

@Component({
  components: {
    SquarePanel,
    UserName,
    TwitterIcon,
    QuestionComments,
  },
})
export default class QuestionPanelDetail extends Vue {

  // これどうやねん
  private fa: boolean = false;
  private tr: boolean = true;

  // 元々 string として良いのでは ??
  private questionId: string | Array<(string | null)> = '';
  // データベース Question 通り
  private question: any = {};
  // 質問者の名前
  private userName: string = '';
  // 質問者の twitter Id;
  private userTwitterId: string = '';
  // ユーザの名前
  private name: string = '';
  private uid: string = '';
  private isReady: boolean = false;
  private alert: boolean = false;

  private isFavorited: boolean = false;
  private isBookMarked: boolean = false;


  private created(): void {
    if (this.getToken() != null) {
      const claims = JSON.parse(atob(this.getToken().split('.')[1]));
      this.uid = claims.uid;
      this.name = claims.name;
    }
    this.questionId = this.$route.query.questionId;
    this.createQuestion();
    this.browseQuestion();
    this.updateIsFavorite();
    this.updateIsBookMarked();
  }

  // ブックマーク状態を更新する
  private updateIsBookMarked(): void {
    const url = 'api/book-mark/' + this.uid + '/' + this.questionId;
    const method = 'GET';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.isBookMarked = json;
    });
  }

  // いいね状態を更新する
  private updateIsFavorite(): void {
    const url = 'api/question-good/' + this.uid + '/' + this.questionId;
    const method = 'GET';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.isFavorited = json;
    });
  }

  private updateQuestionCompleted(): void {
    const url = 'api/question/' + String(this.questionId) + '/completed';
    const method = 'PUT';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      location.reload();
    });
  }

  private favoriteQuestion(): void {
    const url = 'api/question/' + String(this.questionId) + '/favorite';
    const method = 'put';
    const headers = {authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      this.createQuestion();
      this.updateIsFavorite();
    });
  }

  private bookMark(): void {
    const url = 'api/book-mark/' + this.questionId;
    const method = 'put';
    const headers = {authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      this.createQuestion();
      this.updateIsBookMarked();
    });
  }

  private browseQuestion(): void {
    const url = 'api/no-auth/question/' + this.questionId + '/browse';
    const method = 'put';
    fetch(url, {method});

  }

  private createQuestion(): void {
    const url = '/api/no-auth/question/' + this.questionId;

    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.question = json;
      this.isReady = true;
      this.setUser();
    });
  }

  private setUser(): void {

    const url = 'api/no-auth/user/' + this.question.uid;
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {headers}).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.userName = json.name;
      this.userTwitterId = json.twitter_id;
    });
  }

  private deleteQuestion(): void {
    const url = 'api/question/' + this.questionId;
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        // 質問を削除しました
        alert('質問を削除しました');
        this.$router.push('./');
      }
    });
  }

  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>

<style scoped>
.small {
  font-size: 75%;
}
.question-panel-detail {
  margin: 10px;
}
.mavon-editor {
  height: 100;
  z-index: 0;
}
.question-title {
  font-size: 30px;
  letter-spacing: 0.05em;
  color: rgb(66, 66, 66);
}
.url {
  margin-left: 20px;
  font-size: 16px;
}
</style>
