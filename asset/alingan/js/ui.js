class UI {

    constructor() {
        this.transactionItemList = []
        this.grandTotal = 0;
    }

    checkListStatus() {
        if (this.transactionItemList < 1) {
            return false
        } else {
            return true
        }
    }

    openAddProductModal () {
        const addProductToListModal = document.getElementById("add-product-to-list-modal");
        addProductToListModal.classList.add("is-active")
    }

    closeAddProductModal() {
        
        const productNameForm = document.getElementById("add-product-modal-product-name");
        const productPriceForm = document.getElementById("add-product-modal-product-price");
        const productMsuForm = document.getElementById("add-product-modal-msu");
        const subtotalForm = document.getElementById("add-product-modal-sub-total");
        const qtyForm = document.getElementById("add-product-modal-qty");

        productNameForm.value = "";
        productPriceForm.value = "";
        productMsuForm.value = "";
        subtotalForm.value = 0;
        qtyForm.value = 0;

        const addProductToListModal = document.getElementById("add-product-to-list-modal");
        addProductToListModal.classList.remove("is-active")
    }

    addDataToAddProductModal(id, name, price, msu) {
        const productIdForm = document.getElementById("add-product-modal-product-id");
        const productNameForm = document.getElementById("add-product-modal-product-name");
        const productPriceForm = document.getElementById("add-product-modal-product-price");
        const productMsuForm = document.getElementById("add-product-modal-msu");
        productIdForm.value = id;
        productNameForm.value = name;
        productPriceForm.value = price;
        productMsuForm.value = msu;
    }

    calculateSubTotalOnChange() {
        const price = document.getElementById("add-product-modal-product-price").value;
        const qty = document.getElementById("add-product-modal-qty").value;
        document.getElementById("add-product-modal-sub-total").value = parseFloat(price)*parseFloat(qty);
    }

    addProductToTransactionList() {
        
        let item = new TransactionItem(
            document.getElementById("add-product-modal-product-id").value,
            document.getElementById("add-product-modal-product-name").value,
            parseFloat(document.getElementById("add-product-modal-product-price").value),
            parseFloat(document.getElementById("add-product-modal-qty").value),
            parseFloat(document.getElementById("add-product-modal-sub-total").value))
                
        this.transactionItemList = [...this.transactionItemList, item];
        this.reRenderList();
        this.reCountGrandTotal();
        this.closeAddProductModal();

    }

    removeProductFromTransactionList(removedId) {

        let filtered = this.transactionItemList.filter((item)=> {
                if (item.productId != removedId) {
                    return true
                } else {
                    return false
                }
        })

        this.transactionItemList = [...filtered]

        this.reRenderList();
        this.reCountGrandTotal();
    }

    createTransactionListRow(item) {
        
        const td1 = document.createElement("td");
        td1.appendChild(document.createTextNode(item.productId))
        td1.style.cssText = "display: none"

        const td2 = document.createElement("td");
        td2.appendChild(document.createTextNode(item.product))

        const td3= document.createElement("td");
        td3.appendChild(document.createTextNode(item.price))

        const td4 = document.createElement("td");
        td4.appendChild(document.createTextNode(item.quantity))

        const td5 = document.createElement("td");
        td5.appendChild(document.createTextNode(item.subtotal))
        const tr = document.createElement("tr");

        const textRemove = document.createTextNode("remove");
        const buttonRemove = document.createElement("button");
        buttonRemove.classList.add("button");
        buttonRemove.classList.add("is-danger");
        buttonRemove.classList.add("is-small");
        buttonRemove.classList.add("remove-btn");
        buttonRemove.appendChild(textRemove);

        const td6 = document.createElement("td");
        td6.appendChild(buttonRemove);
        
        tr.appendChild(td1);
        tr.appendChild(td2);
        tr.appendChild(td3);
        tr.appendChild(td4);
        tr.appendChild(td5);
        tr.appendChild(td6); 

        return tr;

    }

    reRenderList() {

        const tableList = document.getElementById("transaction-list-table-body");

        while (tableList.lastElementChild) {
            tableList.removeChild(tableList.lastElementChild);
        }
        
        for (var i = 0; i <= this.transactionItemList.length-1;i++) {
            tableList.appendChild(this.createTransactionListRow(this.transactionItemList[i]));
        }
    }

    reCountGrandTotal() {

        this.grandTotal = 0;

        for (var i = 0; i <= this.transactionItemList.length-1;i++) {
            this.grandTotal = this.grandTotal + parseFloat(this.transactionItemList[i].subtotal);
        }

        const grandTotalWrapper = document.getElementById("transaction-list-grand-total-wrapper");
        var total = document.createTextNode(this.grandTotal);
        var p = document.createElement("p");
        p.setAttribute("id", "transaction-grand-total")
        p.appendChild(total);
        
        grandTotalWrapper.replaceChildren(p);
        
    }
    
}