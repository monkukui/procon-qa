<template>
  <div class="home">
    <div class="main-contents">
      <h1 id="title">{{ user }} さん, こんにちは</h1>
      <button @click="logout">LOGOUT</button>
      <Questions />
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Questions from '@/components/Questions.vue';

@Component({
  components: {
    Questions,
  },
})
export default class Home extends Vue {
  private user: string = '名無しのヨッシー';

  private mounted(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.user = claims.name;
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
  private logout(): void {
    localStorage.removeItem('token')
  }
}
</script>

<style scoped>
.home {
}
</style>
