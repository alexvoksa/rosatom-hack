export default `
query a {
    suppliers (limit: 20) {
        ogrn
        name
        short_name
        inn
        kpp
        registered_at
        okpo
        oktmo_code
        oktmo_name
        description
        reputation
        sold_amount
        successful_tenders
        unsuccessful_tenders
        is_innovate
    }
    supplier_blacklist {
        supplier_inn
    }
}
`