const ui = new UI();

const listOfProducts = document.getElementById("list-of-products")
const closeAddProductToListModalButton = document.getElementById("close-add-stock-to-list-modal")
const addProductModalQtyForm = document.getElementById("add-product-modal-qty")
const addProductToListButton = document.getElementById("add-product-to-list")
const transactionItemList = document.getElementById("transaction-list-table")
const saveTransactionButton = document.getElementById("save-transaction")

listOfProducts.addEventListener("click", (e) => {

    if (e.target.classList.contains("choose-btn")) {
        
        const mediaContentChilds = e.target.parentNode.childNodes;
        const productName = mediaContentChilds[1].textContent;
        const productPrice = parseFloat(mediaContentChilds[3].textContent);
        const productMsu = mediaContentChilds[5].textContent;
        const productId = mediaContentChilds[7].textContent;
        ui.addDataToAddProductModal(productId,productName,productPrice, productMsu)
        ui.openAddProductModal();
    }

})

closeAddProductToListModalButton.addEventListener("click", (e) => {
    ui.closeAddProductModal();
})

addProductModalQtyForm.addEventListener("change", (e) => {
    ui.calculateSubTotalOnChange();
})

addProductToListButton.addEventListener("click", (e)=> {
    ui.addProductToTransactionList();
})

transactionItemList.addEventListener("click", (e) => {
    
    if (e.target.classList.contains("remove-btn")) {
        const row = e.target.parentNode.parentNode.childNodes;
        let id = row[0].textContent;
        ui.removeProductFromTransactionList(id);
    }
    
})

saveTransactionButton.addEventListener("click", async (e) => {
    
    if (!ui.checkListStatus()) {
        return alert("please select product first")
    }

    let request = {}
    request.agentId = "agent-001"
    request.storeId = "str-001"
    request.total = parseFloat(document.getElementById("transaction-grand-total").childNodes[0].nodeValue)
    request.transactionItems = []

    ui.transactionItemList.forEach((item)=> {

        let transactionItem = {}
        transactionItem.productId = item.productId
        transactionItem.usedPrice = parseFloat(item.price)
        transactionItem.buyQuantity = parseInt(item.quantity)

        request.transactionItems.push(transactionItem)

    })

    var endpoint = "/owner/new/transaction/submit"

    try {
        
        await axios.post(endpoint, request)

    } catch (error) {
        alert("Error submit transaction"+error)
    }
 
    window.location.href="/owner/transaction"

})

