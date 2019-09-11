<template>
  <div class="question-form">
    <v-card flat>
      <v-snackbar
        v-model="snackbar"
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
                    Do you accept the
                    <a href="javascript:;" @click.stop="terms = true">terms</a>
                    and
                    <a href="javascript:;" @click.stop="conditions = true">conditions?</a>
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
        states: ['AC', 'WA', 'TLE', 'MLE', 'PE', 'a', 'b', 'c', 'd', 'e'],
        conditions: false,
        content: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur sodales ligula in libero. Sed dignissim lacinia nunc.`,
        snackbar: false,
        terms: false,
        defaultForm,
      }
    },

    computed: {
      formIsValid () {
        return (
          this.form.title &&
          this.form.body
        )
      },
    },
    methods: {
      resetForm () {
        this.form = Object.assign({}, this.defaultForm);
        this.$refs.form.reset();
      },
      getToken() {
        return localStorage.getItem('token');
      },
      submit () {
        
        // TODO ここに POST question の api 処理を記述する
        const url = 'api/questions';
        const method = 'POST';
        const headers = {
          'Authorization': `Bearer ${this.getToken()}`,
          'Content-Type': 'application/json; charset=UTF-8'
        }
        const body = JSON.stringify({
          title: this.form.title,
          body: this.form.body,
        });
        
        console.log(body);

        fetch(url, {method, headers, body}).then(response => {
          if(response.ok) {
            return response.json();
          }
        }).then(json => {
          if(typeof json === 'undefined') {
            return;
          }
          this.todos.push(json);
          this.newTodo = '';
        });


        console.log(this.form);
        this.snackbar = true;
        this.resetForm();
      },
    },
  }
</script>

<style scoped>
</style>
