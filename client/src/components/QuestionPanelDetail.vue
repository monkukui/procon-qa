<template>
  <v-card
    max-width="1000"
    class="mx-auto"
  >
    <v-list-item>
      <v-list-item-avatar color="grey">M</v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title class="headline">
          {{ question.title }}
        </v-list-item-title>
        <v-list-item-subtitle>{{ question.user }}monkukui</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
    <v-divider class="mx-4"></v-divider>

    <v-row>
      <v-col> 
        <v-card-text>
          {{ question.body }}
        </v-card-text>
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
        </v-btn>
        <v-btn
          class="tag"
          color="blue-grey lighten-4"
          x-small
        >
          tag (仮)
        </v-btn>
      </v-col>
      <v-col>
        <v-card-text>URL: <a :href="question.url">{{ question.url }}</a></v-card-text>
        <v-card-text>ステータス: {{ question.state }}</v-card-text>
        <v-card-text>質問日時: {{ question.date }}</v-card-text>
        <v-card-text>回答状況: {{ question.completed }}</v-card-text> <!-- solved にしたら？ -->
      </v-col>
    </v-row>

    <v-card-actions>
      <v-btn
        text
        color="deep-purple accent-4"
      >
        Read
      </v-btn>
      <v-btn
        text
        color="deep-purple accent-4"
      >
        Bookmark
      </v-btn>
      <div class="flex-grow-1"></div>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { Component, Prop, Vue, Emit } from 'vue-property-decorator';
import SquarePanel from '@/components/SquarePanel.vue';

@Component({
  components: {
    SquarePanel,
  },
})
export default class QuestionPanel extends Vue {
  // 元々 string として良いのでは ??
  private questionId: number;

  // データベース Question 通り
  private question = {};

  // 質問者の名前
  private userName: string = '';

  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
    this.createQuestion();
  }

  private createQuestion(): void {
    const url = 'api/question/' + String(this.questionId);
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

    const url = 'api/user/' + String(this.question.uid);
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
  margin: 3%;
}

.tag {
  margin-bottom: 1%;
  margin-left: 2%;
}
</style>
