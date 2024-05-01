
export const validation = async (schema, values) => {
    let schemaErrors = null

    try {
        const response = await schema.validate(values, { abortEarly: false })
    } catch (err) {
        let Errors = err.inner?.map(err => {
            return { [err.path]: err.message }
        })

        if (Errors)
            schemaErrors = Object.assign({}, ...Errors)
        else
            schemaErrors = null
    }

    return { ...schemaErrors }
}