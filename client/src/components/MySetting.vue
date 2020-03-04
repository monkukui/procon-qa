<template>
  <div class="myprofile">
    <hr>
    <WithdrawButton />
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import WithdrawButton from '@/components/WithdrawButton.vue';

@Component({
  components: {
    WithdrawButton,
  },
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
