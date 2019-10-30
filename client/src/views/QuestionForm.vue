<template>
  <div class="question-form">
    <v-card flat>
      <v-snackbar v-model="snackbar"
        absolute
        top
        right
        color="success"
      >
        <span>Registration successful!</span>
        <v-icon dark>mdi-checkbox-marked-circle</v-icon>
      </v-snackbar>
      <v-form ref="form" @submit.prevent="submit">
        <v-container fluid>
          <v-row>
            <v-col cols="12" sm="7">
              <v-text-field
                v-model="form.title"
                :rules="rules.title"
                color="blue darken-2"
                label="タイトル"
                required
              ></v-text-field>
            </v-col>

            <v-col cols="12" sm="7">
              <v-textarea
                v-model="form.body"
                :rules="rules.body"
                color="teal"
                required
              >
                <template v-slot:label>
                  <div>
                    本文
                  </div>
                </template>
              </v-textarea>
            </v-col>

            <v-col cols="12" sm="6">
              <v-select
                v-model="form.state"
                :items="states"
                color="pink"
                label="ステータス(任意)"
              ></v-select>
            </v-col>
            <v-col cols="12" sm="6">
              <v-text-field
                v-model="form.url"
                :rules="rules.url"
                color="pink"
                label="URL(任意)"
              ></v-text-field>
            </v-col>

            <v-col cols="12">
              <v-checkbox
                v-model="form.terms"
                color="green"
              >
                <template v-slot:label>
                  <div @click.stop="">
                    <a href="javascript:;" @click.stop="terms = true">誓約書</a>に同意しますか?
                  </div>
                </template>
              </v-checkbox>
            </v-col>
          </v-row>
        </v-container>
        <v-card-actions>
          <div class="flex-grow-1"></div>
          <v-btn
            :disabled="!formIsValid"
            color="primary"
            type="submit"
            x-large
          >送信</v-btn>
        </v-card-actions>
      </v-form>
      <v-dialog v-model="terms" width="70%">
        <v-card>
          <v-card-title class="title">Terms</v-card-title>
          <v-card-text v-for="n in 5" :key="n">
            {{ content }}
          </v-card-text>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn
              text
              color="purple"
              @click="terms = false"
            >Ok</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog v-model="conditions" width="70%">
        <v-card>
          <v-card-title class="title">Conditions</v-card-title>
          <v-card-text v-for="n in 5" :key="n">
            {{ content }}
          </v-card-text>
          <v-card-actions>
            <div class="flex-grow-1"></div>
            <v-btn
              text
              color="purple"
              @click="conditions = false"
            >Ok</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
  </div>
</template>

<script>
// TODO ts にする ... ? 
// 正直 ts にするのめんどくさくなってきた
  export default {
    data () {
      const defaultForm = Object.freeze({
        title: '',
        body: '',
        state: '',
        url: '',
        terms: false,
      })
      return {
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
        defaultForm,
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
          this.form.body &&
          this.form.terms
        )
      },
    },
    methods: {
      resetForm() {
        this.form = Object.assign({}, this.defaultForm);
        this.$refs.form.reset();
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
</style>
