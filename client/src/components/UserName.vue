<template>
  <span>
    <span v-if="isExist">
      <router-link 
        :class="color"
        :to="{ name: 'userpage', query: { uid: this.uid }}"
      >
        <span v-if="size==='large'">
          <h1>{{ name }}</h1>
        </span>
        <span v-else class="name">{{ name }}</span>
        
      </router-link>
    </span>
    <span v-else>
      {{ name }}
    </span>
  </span>
</template>


<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class UserName extends Vue {
  @Prop()
  private uid!: string;
  @Prop()
  private size!: string;

  private color: string = '';

  private name: string = '';
  private isExist: boolean = false;

  private created(): void {
    this.setUser();
  }

  private setUser(): void {
    const url = 'api/no-auth/user/' + this.uid;
    fetch(url).then((response) => {
      if (response.ok) {
        return response.json();
      }
      return [];
    }).then((json) => {
      this.name = json.name;
      const favoriteSum = json.favorite_sum;
      if (this.name === '') {
        this.name = '退会済みのユーザ';
      } else {
        this.isExist = true;
      }


      if (favoriteSum < 1) {
        this.color = 'user-unrated';
      } else if (favoriteSum < 2) {
        this.color = 'user-gray';
      } else if (favoriteSum < 4) {
        this.color = 'user-brown';
      } else if (favoriteSum < 8) {
        this.color = 'user-green';
      } else if (favoriteSum < 16) {
        this.color = 'user-cyan';
      } else if (favoriteSum < 32) {
        this.color = 'user-blue';
      } else if (favoriteSum < 64) {
        this.color = 'user-yellow';
      } else if (favoriteSum < 128) {
        this.color = 'user-orange';
      } else {
        this.color = 'user-red';
      }
    });
  }
}
</script>

<style scoped>
.user-red {
  color:#FF0000;
  text-decoration: none;
}
.user-orange {
  color:#FF8000;
  text-decoration: none;
}
.user-yellow {
  color:#C0C000;
  text-decoration: none;
}
.user-blue {
  color:#0000FF;
  text-decoration: none;
} 
.user-cyan {
  color:#00C0C0;
  text-decoration: none;
}
.user-green {
  color:#008000;
  text-decoration: none;
}
.user-brown {
  color:#804000;
  text-decoration: none;
}
.user-gray {
  color:#808080;
  text-decoration: none;
}
.user-unrated {
  color:#000000;
  text-decoration: none;
}
.user-admin {
  color:#C000C0;
  text-decoration: none;
}
.name {
  font-weight: bold;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-decoration: none;
}
</style>
