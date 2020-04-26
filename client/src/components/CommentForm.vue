<template>
  <div class="comment-form">
    <div v-if="!isLoggedIn">
      <v-alert
      outlined
      type="warning"
      prominent
      border="left"
    >
    <v-row>
      <v-col cols="12" sm="8">
        ログインする必要があります.
      </v-col>
      <v-col cols="12" sm="4">
        <v-btn large color="primary" to="/login">ログイン</v-btn>
      </v-col>
    </v-row>
    </v-alert>
    </div>
    <div v-else>
      <v-textarea
        v-model="text"
      />
      <v-btn
        color="primary"
        outlined
        @click="answer"
        small
        :disabled="text == ''"
      >コメントを送信</v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class CommentForm extends Vue {
  @Prop()
  private qid!: string;
  @Prop()
  private aid!: string;

  private text: string = '';
  private isLoggedIn: boolean = false;
  private created(): void {
    this.isLoggedIn = (this.getToken() != null);
  }
  private getToken(): string | null {
    return localStorage.getItem('token');
  }

  private answer(): void {
    if (this.getToken() === null) {
      alert('server error');
      return;
    }
    const url = (this.qid !== '-1') ? '/api/question-comment' : 'api/answer-comment';
    const method = 'POST';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };

    const body = (this.qid !== '-1') ? JSON.stringify({
        body: this.text,
        qid: Number(this.qid),
      }) : JSON.stringify({
        body: this.text,
        aid: Number(this.aid),
      });

    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      if (typeof json !== 'undefined') {
        // 親コンポーネントに発火させる
        alert('コメントを送信しました');
      }
    });
  }
}
</script>

<style scoped>
.comment-form {
  margin: 10px;
}
</style>
