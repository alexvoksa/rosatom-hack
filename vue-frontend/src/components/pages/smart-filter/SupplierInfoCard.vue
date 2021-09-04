<template>
  <div class="supplier" :class="{ blacklisted: supplier.blacklisted }">
    <div v-if="supplier.blacklisted" class="extra-info">
      <i class="el-icon-info"></i>
      поставщик в черном списке
    </div>
    <div class="supplier-cell">
      <div class="supplier-info">
        <router-link
          class="supplier-name"
          :to="{
            name: 'SupplierPage',
            params: { id: this.supplier.ogrn },
          }"
        >
          <span>{{ supplier.name }}</span>
          <i class="el-icon-link"></i
        ></router-link>
      </div>
      <div class="rep-log">
        <div class="rep-item reputation">
          репутация {{ supplier.reputation }}
        </div>
        <div class="rep-item successful-tender">
          успешных тендеров {{ supplier.successful_tenders }}
        </div>
        <div class="rep-item unsuccessful-tender">
          невыполненых тендеров {{ supplier.unsuccessful_tenders }}
        </div>
      </div>
    </div>
    <div class="supplier-cell">
      <div>
        процент успешных тендеров:
        {{
          Math.floor(
            supplier.successful_tenders /
              (supplier.successful_tenders + supplier.unsuccessful_tenders)
          ) || "0/0"
        }}
      </div>
    </div>
    <div class="supplier-cell"></div>
  </div>
</template>

<script>
/**
 * 
 *  ogrn
    name
    short_name
    email
    phone
    inn
    kpp
 * 
 */

export default {
  name: "SupplierInfoCard",
  props: {
    ind: {
      type: Number,
      required: true,
    },
    supplier: {
      type: Object,
      required: true,
    },
  },
  methods: {
    goToSupplierPage() {
      this.$router.push({
        name: "SupplierPage",
        params: { id: this.supplier.ogrn },
      });
    },
  },
};
</script>

<style scoped>
.title {
  display: flex;
}

.supplier {
  position: relative;
  display: flex;
  flex-direction: row;
  align-items: center;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  border-radius: 3px;
  margin-bottom: 32px;
}

.supplier-cell {
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.supplier-info {
  display: flex;
  align-items: center;
}

.supplier-name {
  font-size: 13px;
  padding: 10px;
  border-radius: 3px;
  background-color: rgba(205, 215, 245, 0.5);
  color: rgb(36, 124, 255);
}

.supplier-name:hover {
  background-color: rgba(205, 215, 245, 0.8);
  color: rgb(0, 100, 250);
  cursor: pointer;
  text-decoration: underline;
}

.rep-log {
  margin-top: 10px;
  display: flex;
}

.rep-item {
  margin-right: 10px;
}

.reputation {
  font-size: 13px;
  padding: 10px;
  border-radius: 3px;
  background-color: rgba(245, 229, 205, 0.7);
  color: rgb(68, 7, 7);
}

.successful-tender {
  font-size: 13px;
  padding: 10px;
  border-radius: 3px;
  background-color: rgba(205, 245, 214, 0.7);
  color: rgb(7, 68, 40);
}

.unsuccessful-tender {
  font-size: 13px;
  padding: 10px;
  border-radius: 3px;
  background-color: rgba(245, 210, 205, 0.7);
  color: rgb(68, 21, 7);
}

.blacklisted {
  background-color: rgba(255, 0, 0, 0.13);
}

.extra-info {
  position: absolute;
  top: -20px;
}
</style>