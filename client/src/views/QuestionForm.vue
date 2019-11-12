<template>
  <div class="question-form">
    <v-card flat>
      <v-snackbar v-model="snackbar"
        absolute
        top
        right
        color="success"
      >
        <span>質問を投稿しました!!</span>
        <v-icon dark>mdi-checkbox-marked-circle</v-icon>
      </v-snackbar>
      <v-form ref="form" @submit.prevent="submit">
        <v-container fluid>
          <v-row>
            <v-col cols="12" sm="8">
              <v-text-field
                v-model="form.title"
                @mouseover="focusTitle"
                @focus="focusTitle"
                :rules="rules.title"
                color="blue darken-2"
                label="タイトル"
                required
                single-line
                outlined
              ></v-text-field>
              <div class="mavon-editor" @mouseover="focusBody">
                <mavon-editor
                  v-model="form.body"
                  language="en"
                  :toolbars="markdownOption"
                />
              </div>

              <v-text-field
                v-model="form.url"
                @mouseover="focusUrl"
                @focus="focusUrl"
                :rules="rules.url"
                color="blue darken-2"
                label="URL(任意)"
                single-line
                outlined
              ></v-text-field>

              <v-btn
                :disabled="!formIsValid"
                color="primary"
                type="submit"
                x-large
              >送信</v-btn>
            </v-col>
            <v-col>
              <div v-if="focus=='title'">
                <TitleTips />
              </div>
              <div v-if="focus=='body'">
                <BodyTips />
              </div>
              <div v-if="focus=='url'">
                <UrlTips />
              </div>
            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </v-card>
  </div>
</template>

<script>
// TODO ts にする ... ? 
// 正直 ts にするのめんどくさくなってきた
import TitleTips from '@/components/TitleTips.vue';
import BodyTips from '@/components/BodyTips.vue';
import UrlTips from '@/components/UrlTips.vue';

export default {
  components: {
    TitleTips,
    BodyTips,
    UrlTips,
  },
  data () {
    const defaultForm = Object.freeze({
      title: '',
      body: '',
      state: '',
      url: '',
    })
    return {
      focus: 'title',
      form: Object.assign({}, defaultForm),
      rules: {
        title: [val => (val || '').length > 0 || 'この項目は必須です'],
        body: [val => (val || '').length > 0 || 'この項目は必須です'],
      },
      states: ['CE', 'MLE', 'TLE', 'RE', 'OLE', 'IE', 'WA', 'AC', 'WJ', 'WR'],
      conditions: false,
      content: `誓約書`,
      snackbar: false,
      terms: false,
      markdownOption: {
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
      },
    }
  },
  computed: {
    getDate () {
      let now = new Date();
      var Year = String(now.getFullYear());
      var Month = String(now.getMonth()+1);
      var Day = String(now.getDate());
      var Hour = String(now.getHours());
      var Min = String(now.getMinutes());
      var Sec = String(now.getSeconds());
      if(Hour.length == 1) Hour = '0' + Hour;
      if(Min.length == 1) Min = '0' + Hour;
      if(Sec.length == 1) Sec = '0' + Hour;
      return Year + '年' + Month + '月' + Day + '日' + Hour + ':' + Min + ':' + Sec;
    },
    formIsValid () {
      return (
        this.form.title &&
        this.form.body
      )
    },
  },
  methods: {
    resetForm() {
      this.form = Object.assign({}, this.defaultForm);
      this.$refs.form.reset();
    },
    focusTitle() {
      this.focus = 'title';
    },
    focusBody() {
      this.focus = 'body';
    },
    focusUrl() {
      this.focus = 'url';
    },
    getToken() {
      return localStorage.getItem('token');
    },
    submit () {
      
      const url = 'api/questions';
      const method = 'POST';
      const headers = {
        'Authorization': `Bearer ${this.getToken()}`,
        'Content-Type': 'application/json; charset=UTF-8'
      }
      const body = JSON.stringify({
        title: this.form.title,
        body: this.form.body,
        date: this.getDate,
        state: this.form.state,
        url: this.form.url,
      });
      
      fetch(url, {method, headers, body}).then(response => {
        if(response.ok) {
          return response.json();
        }
      }).then(json => {
        if(typeof json === 'undefined') {
          return;
        }
      });
      this.snackbar = true;
      this.resetForm();
    },
  },
}
</script>

<style scoped>
.mavon-editor {
  z-index: 0;
}
</style>
