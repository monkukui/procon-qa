<template>
  <div class="about">

    <h2>競技プログラミング専用の Q&A サイトです</h2>
    <v-divider />
    <br>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> <strong>質問</strong>を投稿することができます</p>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> <strong>回答</strong>を投稿することができます</p>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 質問や回答に<strong>いいね</strong>をすることができます</p>
    <br>
    <br>

    <h2>いいねの総数に応じて，色が付与されます</h2>
    <v-divider />
    <br>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 取得したいいねの総数に応じて，色が付与されます</p>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 質問と回答には，いいねがつきます</p>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 自分の投稿には，いいねをつけることができません</p>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 自分の投稿には，いいねをつけることができません</p>
    <br>
    <br>
    
    <v-data-table
      :headers="headers"
      :items="desserts"
      class="elevation-1"
    >
      <template v-slot:item.calories="{ item }">
        <font :color="getColor(item.calories)">{{ item.calories }}</font>
      </template>
    </v-data-table>

    <br>
    <br>

    <h2>どんな質問でも歓迎します</h2>
    <v-divider />
    <br>
    <p><v-icon large color="rgb(129, 180, 111)">mdi-check</v-icon> 競技プログラミングに関する質問なら，なんでも ok</p>
      <ul>
        <li>AtCoder ABC123 の C 問題で，自分の解答コードが WA になるのはなぜ！？</li>
        <li><strong>動的計画法</strong>の漸化式を思いつくまでの過程が知りたいです</li>
        <li><strong>ダイクストラ法</strong>のオススメの問題を教えてください</li>
        <li>区間 max 更新，区間和取得は<strong>遅延セグ木</strong>に乗りますか？</li>
        <li>など</li>
      </ul>
    <br>
    <br>

    <p>ご意見ご要望は
      <v-btn icon href="https://twitter.com/monkukui2">
        <v-icon color="rgb(73, 160, 236)">mdi-twitter</v-icon>
      </v-btn>
      <v-btn icon href="https://github.com/monkukui/procon-qa">
        <v-icon color="rgb(37, 41, 46)">mdi-github</v-icon>
      </v-btn>
      まで．
    </p>
  </div>
</template>
<script>
  export default {
    data () {
      return {
        headers: [
          {
            text: 'いいねの総数',
            align: 'start',
            value: 'name',
          },
          { text: '色', value: 'calories' },
        ],
        desserts: [
          {
            name: '[128, INF)',
            calories: '赤色',
          },
          {
            name: '[64, 128)',
            calories: '橙色',
          },
          {
            name: '[32, 64)',
            calories: '黄色',
          },
          {
            name: '[16, 32)',
            calories: '青色',
          },
          {
            name: '[8, 16)',
            calories: '水色',
          },
          {
            name: '[4, 8)',
            calories: '緑色',
          },
          {
            name: '[2, 4)',
            calories: '茶色',
          },
          {
            name: '[1, 2)',
            calories: '灰色',
          },
          {
            name: '[0, 1)',
            calories: '黒色',
          },
        ],
      }
    },
    methods: {
      getColor (calories) {
        if (calories == '赤色') return 'red';
        else if (calories == '橙色') return 'orange';
        else if (calories == '黄色') return 'yellow';
        else if (calories == '青色') return 'blue';
        else if (calories == '水色') return 'cyan';
        else if (calories == '緑色') return 'green';
        else if (calories == '茶色') return 'brown';
        else if (calories == '灰色') return 'grey';
        else return 'black';
      },
    },
  }
</script>


<style scoped>
.about {
  margin: 50px;
}
</style>
