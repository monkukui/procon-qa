<template>
  <div class="myprofile">
    <hr>
    <p>{{ userName }} さん，こんにちは．</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component({
})
export default class MySetting extends Vue {

  private userName: string = '';
  private created(): void {
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.userName = claims.name;
  }
  private getToken(): any {
    return localStorage.getItem('token');
  }
}
</script>
