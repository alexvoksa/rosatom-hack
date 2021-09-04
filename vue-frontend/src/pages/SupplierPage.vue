<template>
  <div class="supplier">
    <el-dialog
      :title="`Письмо для ${supplier.name}`"
      :visible.sync="emailDialogVisible"
    >
      <div class="email-dialog">
        <div>Тема письма</div>
        <el-input
          v-model="emailTheme"
          placeholder="Тема письма"
          label="Тема письма"
          class="email-theme"
        >
        </el-input>
        <el-input
          v-model="emailText"
          placeholder="Текст письма..."
          type="textarea"
          height="300"
        >
        </el-input>
        <el-select v-model="email" placeholder="Контактная почта">
          <el-option
            v-for="email in emails"
            :key="email"
            :label="email"
            :value="email"
          >
          </el-option>
        </el-select>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="emailDialogVisible = false">Отмена</el-button>
        <el-button type="success" @click="sendEmail">Отправить</el-button>
      </span>
    </el-dialog>

    <router-link :to="{ name: 'SmartFilter' }" class="navigation"
      >к фильтрации</router-link
    >
    <h2>{{ supplier.name }}</h2>
    <div class="info-container">
      <el-button
        type="success"
        class="dialog-toggler"
        @click="emailDialogVisible = true"
        >Написать поставщику <i class="el-icon-edit"></i
      ></el-button>
      <div class="tenders">
        <div v-for="tender in supplier.tenders" :key="tender.id" class="tender">
          tender info
        </div>
      </div>
      <div class="unique-info">
        <div class="economics">
          <div class="economics-item inn">ИНН {{ supplier.inn }}</div>
          <div class="economics-item kpp">КПП {{ supplier.kpp }}</div>
          <div class="economics-item okpo">ОКПО {{ supplier.okpo }}</div>

          <div class="economics-item oktmo_code">
            ОКТМО {{ supplier.oktmo_code }}
          </div>
          <div class="economics-item oktmo_name">
            {{ supplier.oktmo_name }}
          </div>
        </div>
        <div class="contacts">
          <div
            class="contact"
            v-for="(contact, ind) in supplier.supplier_contacts"
            :key="contact.supplier_id + '-' + ind"
          >
            <div v-if="ind === 0" class="contacts-header">Контакты</div>
            <div class="phone">
              <i class="el-icon-phone"></i>
              {{ contact.phone }}
            </div>
            <div class="email">{{ contact.email }}</div>
            <div class="contact-person">
              {{
                `${contact.first_name} ${contact.middle_name} ${contact.last_name}`
              }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/**
 * тендеры, детали тендеров
 *
 * контакты
 *
 * блок инн, кпп
 */
export default {
  name: "SupplierPage",
  data() {
    return {
      supplier: {},
      email: "",
      emails: [],
      emailText: "",
      emailTheme: "Тендер",
      emailDialogVisible: false,
    };
  },
  created() {
    this.$http.hasura
      .post("/", {
        query: require("@/gql/qFetchSupplier.js").default(
          this.$route.params.id
        ),
      })
      .then((resp) => {
        if (resp.data.errors) {
          return;
        }
        this.supplier = resp.data.data.suppliers[0];
        this.emails = this.supplier.supplier_contacts
          .filter((c) => !!c.email)
          .map((c) => c.email);

        this.email = this.emails[0];
      });
  },
  methods: {
    sendEmail() {
      this.$http.fastApi
        .post("/email/send", {
          to: this.email,
          subject: this.emailTheme,
          text: this.emailText,
        })
        .then((resp) => {
          resp;
          this.emailDialogVisible = false;
        });
    },
  },
};
</script>

<style scoped>
.navigation {
  position: absolute;
  top: 120px;
  left: 50px;
}

.email-dialog {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.email-theme {
  margin-top: 10px;
  margin-bottom: 20px;
}

.supplier {
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-top: 50px;
  width: 80%;
}

.dialog-toggler {
  position: absolute;
  left: -6%;
}

.info-container {
  position: relative;
  display: flex;
  justify-content: center;
  width: 100%;
  margin-top: 50px;
}

.tenders {
  width: 60%;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  border-radius: 3px;
  margin-right: 30px;
}

.unique-info {
  display: flex;
  flex-direction: column;
}

.economics {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  border-radius: 3px;
}

.economics-item {
  padding: 2px;
}

.contacts {
  margin-top: 30px;
}

.contacts-header {
  font-weight: bold;
  padding-top: 5px;
  padding-bottom: 10px;
}

.contact {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  border-radius: 3px;
  padding: 10px;
}
</style>