<template>
  <div class="question-panel-detail">
    <v-card
      class="mx-auto"
    >
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="headline">
            {{ question.title }}
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-divider class="mx-4"></v-divider>

      <v-row>
        <v-col md="8"> 
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
          <v-card-text
            v-if="question.url"
          >
            URL: <a :href="question.url">{{ question.url }}</a>
          </v-card-text>
          <!--v-btn
            class="tag"
            color="blue-grey lighten-4"
            x-small
          >
            tag (仮)
          </v-btn>
          <v-btn
            class="tag"
            color="blue-grey lighten-4"
            x-small
          >
            tag (仮)
          </v-btn>
          <v-btn
            class="tag"
            color="blue-grey lighten-4"
            x-small
          >
            tag (仮)
          </v-btn-->
        </v-col>
        <v-col md="4">
          <v-card
            class="mx-auto"
            max-width="300"
            outlined
            color="teal lighten-5"
          >
            <v-list-item three-line>
              <v-list-item-content>
                <div>質問者</div>
                <v-list-item-title class="headline mb-1">{{ userName }}</v-list-item-title>
              </v-list-item-content>
              <v-list-item-avatar
                circle
                size="80"
                color="grey"
              ></v-list-item-avatar>
            </v-list-item>
          </v-card>
        </v-col>
      </v-row>
      <v-chip
        class="ma-2"
        color="pink"
        text-color="white"
        @click="favoriteQuestion"
      >
        <v-avatar
          left
          class="pink darken-4"
        >
        {{ question.favoriteCount }}
        </v-avatar>
        いいね
      </v-chip>
      <v-card-actions>
        <v-card-text>投稿日時: {{ question.date }}</v-card-text>
        <v-btn icon>
          <v-icon @click="favoriteQuestion">mdi-heart</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon @click="updateQuestionCompleted">mdi-share-variant</v-icon>
        </v-btn> 
        <v-btn icon>
          <v-icon @click="deleteQuestion">mdi-home</v-icon>
        </v-btn> 
      </v-card-actions>
    </v-card>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';

@Component({
  components: {
    SquarePanel,
  },
})
export default class QuestionPanelDetail extends Vue {

  // これどうやねん
  private fa: boolean = false;
  private tr: boolean = true;

  // 元々 string として良いのでは ??
  private questionId!: number;

  // データベース Question 通り
  private question: any = {};

  // 質問者の名前
  private userName: string = '';

  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
    this.createQuestion();
  }

  private updateQuestionCompleted(): void {
    const url = 'api/question/' + String(this.questionId) + '/completed';
    const method = 'PUT';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers});
  }

  private favoriteQuestion(): void {
    const url = 'api/question/' + String(this.questionId) + '/favorite';
    const method = 'put';
    const headers = {authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers}).then(() => {
      this.createQuestion();
    });
  }

  private browseQuestion(): void {
    const url = 'api/no-auth/question/' + String(this.questionId) + '/browse';
    const method = 'put';
    const headers = {Authorization: `Bearer ${this.getToken()}`};
    fetch(url, {method, headers});
  }

  private createQuestion(): void {
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
      this.browseQuestion();
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
    });
  }

  private deleteQuestion(): void {
    const url = 'api/question/' + String(this.questionId);
    const method = 'DELETE';
    const headers = {Authorization: `Bearer ${this.getToken()}`};

    fetch(url, {method, headers}).then((response) => {
      if (response.ok) {
        // 質問を削除しました
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
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}
.mavon-editor {
  z-index: 0;
}
</style>
