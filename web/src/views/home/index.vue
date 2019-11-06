<template>
  <el-row class="home">
    <el-row>
      <el-menu default-active="1" class="menu" mode="horizontal" active-text-color="white" text-color="whitesmoke" background-color="transparent">
        <span index="1" class="logo">DC Fortune</span>
        <el-menu-item index="6" class="right"><a href="#register">Register</a></el-menu-item>
        <el-menu-item index="5" class="right"><a href="#prize">Prize</a></el-menu-item>
        <el-menu-item index="3" class="right"><a href="#info">Info</a></el-menu-item>
        <el-menu-item index="1" class="right"><a>Home</a></el-menu-item>
      </el-menu>
    </el-row>
    <el-row class="welcome" type="flex" justify="center">
      <el-col class="left" :span="14">
        <a name="home"></a>
        <div class="title">DC14 Hackathon</div>
        <div class="content">The first Hackathon happen in DC14 with total prizes more than <span class="high-light">10.000.000 VND</span></div>
        <div class="content">From 21 Dec - 22 Dec 2019</div>
        <div class="content">At Floor 4 - Lab 6 - TMA Solutions</div>
        <br>
        <el-button type="warning" style="background-color:#F7B933; color:black;">Register Now</el-button>
      </el-col>
    </el-row>
    <el-row name="info" type="flex" justify="center" class="info" style="background-color:whitesmoke">
      <a name="info"></a>
      <a name="schedule"></a>
      <el-col :span="10" class="column">
        <h2>Topic</h2>
        <p>
            The world is now flat and software also need to change with the speed of the industry 4.0, building a system that can server diversity of users is a big challenge.
            Our topic is to build a chat application that serve multiple users at the same time.
        </p>
      </el-col>
      <el-col :span="10" class="column">
        <h2>Schedule</h2>
        <div>09:00 AM - 09:30 AM: Welcome and broadcast the topic.</div>
        <div>09:30 AM - 10:00 AM: Clarify requirements.</div>
      </el-col>
    </el-row>
    <el-row style="background-color:white">
      <a name="prize"></a>
      <h2>Prize</h2>
    </el-row>
    <el-row style="background-color:whitesmoke">
      <a name="register"></a>
      <h2>Register</h2>
    </el-row>
    <el-row class="footer">
      Copyright © 2019 by Thanh Pham - All Rights Reserved.
    </el-row>
  </el-row>
</template>

<style lang="scss" scoped>
.title {
  font-size: 2em;
}

.section-title {
  font-size: 1.5em;
}

.el-menu.el-menu--horizontal {
  border: none;
}

.high-light {
  font-size: 1.5em;
  color: #F7B933;
}

.home {
  background-image: url('http://hackbu.org/2017s/img/header-background.jpg');
  background-repeat: no-repeat, repeat;

  .logo {
    font-size: 1.3em;
    font-weight: bold;
    color: #F7B933;
    margin-left: 20px;
  }

  .menu {
    top: 0px;

    .right {
      float: right;
    }
  }

  .welcome {
    min-height: 100vh;
    top: 15vh;

    .left {
      text-align: left;
      color: white;
      font-size: 1.5em;
      line-height: 1.75em;

      .content {
        font-size: 0.7em;
        font-style: italic;
        color: white;
        font-family: 'Open Sans', sans-serif;
      }
    }

    .right {
      color: white;
      vertical-align: middle;
    }
  }

  .info {
    text-align: left;
    .column {
        margin-left: 20px;
    }
  }

  .footer {
    top: 20px;
    color: white;
    text-align: center;
    min-height: 100px;
  }

}
</style>
