const ui = new UI();

const listOfProducts = document.getElementById("list-of-products")
const closeAddProductToListModalButton = document.getElementById("close-add-stock-to-list-modal")
const addProductModalQtyForm = document.getElementById("add-product-modal-qty")
const addProductToListButton = document.getElementById("add-product-to-list")
const transactionItemList = document.getElementById("transaction-list-table")
const saveTransactionButton = document.getElementById("save-transaction")

listOfProducts.addEventListener("click", (e) => {

    if (e.target.classList.contains("choose-btn")) {

        const cardHeader = e.target.parentNode.parentNode.parentNode.childNodes[1];
        const cardHeaderChilds = cardHeader.childNodes;
        const productName = cardHeaderChilds[1].innerText;
    
        const cardContent = e.target.parentNode.parentNode.parentNode.childNodes[5];
        const cardContentChilds = cardContent.childNodes;
        const productPrice = parseFloat(cardContentChilds[1].childNodes[1].textContent);
        const productMsu = cardContentChilds[3].childNodes[1].textContent;
        const productId = cardContentChilds[5].textContent;
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
    let source = document.getElementById("pos-input-source").value;
   
    
    if (source === "agent") {
        request.storeId = document.getElementById("pos-input-store-id").value;
        request.agentId = document.getElementById("pos-input-agent-id").value;
    } else {
        var selectStore = document.getElementById("pos-input-store-id");
        request.storeId = selectStore.options[selectStore.selectedIndex].value;
        var selectAgent = document.getElementById("pos-input-agent-id");
        request.agentId = selectAgent.options[selectAgent.selectedIndex].value;
    }

    request.total = parseFloat(document.getElementById("transaction-grand-total").childNodes[0].nodeValue)
    request.transactionItems = []

    ui.transactionItemList.forEach((item)=> {

        let transactionItem = {}
        transactionItem.productId = item.productId
        transactionItem.usedPrice = parseFloat(item.price)
        transactionItem.buyQuantity = parseInt(item.quantity)

        request.transactionItems.push(transactionItem)

    })

    console.log(request)

    var endpoint = "/owner/new/transaction/submit"

    try {
        
        await axios.post(endpoint, request)

    } catch (error) {
        alert("Error submit transaction"+error)
    }
 
    if (source === "agent") {
        window.location.href = "/agent/new/transaction"
    } else {
        window.location.href="/owner/transaction"
    }

})

