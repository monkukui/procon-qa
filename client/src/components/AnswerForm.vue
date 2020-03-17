<template>
  <div class="answer-form">
    <div v-if="!isLoggedIn">
      <v-alert
      outlined
      type="warning"
      prominent
      border="left"
    >
    <v-row>
      <v-col cols="12" sm="10">
        あなたの知見を共有してください.<br>
        質問に回答するには，ログインする必要があります.
      </v-col>
      <v-col cols="12" sm="2">
        <v-btn large color="primary" to="/login">ログイン</v-btn>
      </v-col>
    </v-row>
    </v-alert>
    </div>
    <div v-else>
      <span class="font-weight-black headline">回答</span>
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
        large
        :disabled="text == ''"
      >回答する</v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
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
  private fa: boolean = false;
  private tr: boolean = true;
  private questionId!: number;
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
      favoriteCount: 0,
      date: this.getDate(),
      qid: this.questionId,
    });
    fetch(url, {method, headers, body}).then((response) => {
      if (response.ok) {
        return response.json();
      }
    }).then((json) => {
      if (typeof json !== 'undefined') {
        location.reload();
      }
    });
  }
  private getDate(): string {
    const now = new Date();
    const Year = String(now.getFullYear());
    const Month = String(now.getMonth() + 1);
    const Day = String(now.getDate());
    let Hour = String(now.getHours());
    let Min = String(now.getMinutes());
    let Sec = String(now.getSeconds());
    if (Hour.length === 1) {
      Hour = '0' + Hour;
    }
    if (Min.length === 1) {
      Min = '0' + Hour;
    }
    if (Sec.length === 1) {
      Sec = '0' + Hour;
    }
    return Year + '年' + Month + '月' + Day + '日' + Hour + ':' + Min + ':' + Sec;
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
