export default (sid) => `
query fetchSupplier {
    suppliers (where: {
        ogrn: {_eq: ${sid}}
    }) {
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

        supplier_contacts {
            phone
            email
            address
            first_name 
            middle_name 
            last_name
            description
            supplier_id 
        }
    }
}
`