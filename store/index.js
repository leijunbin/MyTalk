export const state = () => ({
    AccountDialog: false,
    LeftNavigation: false,
})

export const mutations = {
    SET_ACCOUNT_DIALOG(state) {
        state.AccountDialog = !state.AccountDialog
    },
    SET_SIDE_STATUS(state, val) {
        state.LeftNavigation = val
    }
}

export const actions = {
    
}