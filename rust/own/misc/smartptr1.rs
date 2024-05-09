use std::rc::Rc;
use std::cell::RefCell;

struct Cat {
    name: String,
    parent: Option<Rc<RefCell<Cat>>>
}

impl Drop for Cat {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data {}", self.name)
    }
}

fn main() {
    bm();

    let cat1 = Rc::new(RefCell::new(
    Cat { name: "Vasja".to_string(), parent: None }
    ));
    // let cat2 = Rc::new(RefCell::new(Cat { name: "Alisa".to_string(), parent: None }));
    let cat3 = Rc::new(RefCell::new(Cat {
        name: "Rizhik".to_string(),
        parent: Some(Rc::clone(&cat1))   // склонировал умный указатель, но не данные
    }));
    let cat4 = Rc::new(RefCell::new( Cat {
        name: "Kuzja".to_string(),
        parent: Some(Rc::clone(&cat3))
    }));
    println!("Strong count cat1: {}", Rc::strong_count(&cat1)); // 2
    cat3.borrow_mut().name = "Rizhik II".to_string();

    // Допустимо изменять родителя cat1 через ребёнка cat3:
    cat3.borrow().parent.as_ref().unwrap().borrow_mut().name = "Vaska III".to_string();

    parents_iterate(&cat4);

    std::mem::drop(cat3);
    println!("End, csp1.data: {}", cat1.borrow().name);
}

fn parents_iterate(child: &Rc<RefCell<Cat>>) {
    let mut current = Some(Rc::clone(child));
    while let Some(current_csp) = current.take() {
        let current_csp_ref = current_csp.borrow();
        println!("Getting csp data: {}", current_csp_ref.name);
        if let Some(parent) = current_csp_ref.parent.as_ref() {
            let parent_csp_ref = parent.borrow();
            println!("Getting parent csp data: {}", parent_csp_ref.name);

            current = Some(Rc::clone(parent));
        }

        println!("=======");
    }
}

fn bm() {
    let a = RefCell::new(5);
    let b = &a;
    *a.borrow_mut() += 1;
    println!("a: {a:?}, b: {b:?}");
}