<template>
  <div class="answer-form">
    <v-snackbar
      v-model="snackbar"
      right
      color="success"
    >
      <strong>{{ snackbarText }}</strong>

      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <div v-if="!isLoggedIn">
      <v-alert
        outlined
        type="warning"
        prominent
        border="left"
      >
    <v-row>
      <v-col cols="12" sm="8">
        あなたの知見を共有してください.<br>
        質問に回答するには，ログインする必要があります.
      </v-col>
      <v-col cols="12" sm="4">
        <v-btn large color="primary" to="/login">ログイン</v-btn>
      </v-col>
    </v-row>
    </v-alert>
    </div>
    <div v-else>
      <br>
      <div class="mavon-editor">
        <mavon-editor
          :toolbars="markdownOption"
          v-model="text"
          language="ja"
          placeholder='回答本文'
          :boxShadow="fa"
          :ishljs="fa"
        />
      </div>
      <br>
      <v-btn
        color="primary"
        @click="answer"
        x-large
        :disabled="text == ''"
      >送信</v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from 'vue-property-decorator';
@Component({
  components: {
  },
})
export default class AnswerForm extends Vue {
  private markdownOption: any = {
    bold: true,
    italic: true,
    header: true,
    underline: true,
    strikethrough: true,
    mark: true,
    quote: true,
    ol: true,
    ul: true,
    code: true,
    table: true,
    help: true,
    alignleft: true,
    aligncenter: true,
    alignright: true,
    subfield: true,
    preview: true,
    // false
    link: false,
    imagelink: false,
    superscript: false,
    subscript: false,
    undo: false,
    redo: false,
    fullscreen: false,
    readmodel: false,
    htmlcode: false,
    trash: false,
    save: false,
    navigation: false,
  };
  @Prop()
  private uid!: string;
  private fa: boolean = false;
  private tr: boolean = true;
  private snackbar: boolean = false;
  private questionId!: number;
  private snackbarText: string = "回答を送信しました";
  private text: string = '';
  private isLoggedIn: boolean = false;
  private created(): void {
    this.questionId = Number(this.$route.query.questionId);
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
    const url = 'api/answers';
    const method = 'POST';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };
    const body = JSON.stringify({
      body: this.text,
      qid: this.questionId,
    });
    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        // 親コンポーネントに発火させる
        this.postNotification();
        this.$emit('answer');
      }
    });
  }

  // 通知を発行
  private postNotification(): void {
    const url = 'api/notification/' + this.uid + '/' + this.questionId + '/1';
    const method = 'POST';
    
    const token = this.getToken();
    if (token === null) {
      alert('server error');
      return;
    }

    const claims = JSON.parse(atob(token.split('.')[1]));

    // コメントする側の id が uid と同じなら，通知を発行しない
    if (this.uid === claims.uid) {
      this.snackbar = true;
      this.text = '';
      return;
    }

    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };

    // body もつけるか TODO
    const body = JSON.stringify({
      body: this.text,
    });

    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        this.userUpdate();
        this.snackbar = true;
        this.text = '';
      }
    });
  }

  // uid さんの NotficationFlag を true にする
  private userUpdate(): void {
    const url = '/api/notification/' + this.uid + '/1';
    const method = 'PUT';
    const headers = {
      'Authorization': `Bearer ${this.getToken()}`,
      'Content-Type': 'application/json; charset=UTF-8',
    };
    fetch(url, {method, headers});
  }
}
</script>

<style scoped>
.answer-form {
  margin: 10px;
}
.mavon-editor {
  z-index: 0;
}
</style>
